package frontend

import (
	"context"
	"runtime"
	"sync"
	"time"

	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/api/serviceerror"
	versionpb "go.temporal.io/api/version/v1"
	"go.temporal.io/server/common/headers"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/persistence"
	"go.temporal.io/server/common/primitives/timestamp"
	"go.temporal.io/server/common/rpc/interceptor"
	"go.temporal.io/version/check"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const VersionCheckInterval = 24 * time.Hour

type VersionChecker struct {
	config                 *Config
	shutdownChan           chan struct{}
	metricsHandler         metrics.Handler
	clusterMetadataManager persistence.ClusterMetadataManager
	startOnce              sync.Once
	stopOnce               sync.Once
	sdkVersionRecorder     *interceptor.SDKVersionInterceptor
}

func NewVersionChecker(
	config *Config,
	metricsHandler metrics.Handler,
	clusterMetadataManager persistence.ClusterMetadataManager,
	sdkVersionRecorder *interceptor.SDKVersionInterceptor,
) *VersionChecker {
	return &VersionChecker{
		config:                 config,
		shutdownChan:           make(chan struct{}),
		metricsHandler:         metricsHandler.WithTags(metrics.OperationTag(metrics.VersionCheckScope)),
		clusterMetadataManager: clusterMetadataManager,
		sdkVersionRecorder:     sdkVersionRecorder,
	}
}

func (vc *VersionChecker) Start() {
	if vc.config.EnableServerVersionCheck() {
		vc.startOnce.Do(func() {
			// TODO: specify a timeout for the context
			ctx := headers.SetCallerInfo(
				context.TODO(),
				headers.SystemBackgroundHighCallerInfo,
			)

			go vc.versionCheckLoop(ctx)
		})
	}
}

func (vc *VersionChecker) Stop() {
	if vc.config.EnableServerVersionCheck() {
		vc.stopOnce.Do(func() {
			close(vc.shutdownChan)
		})
	}
}

func (vc *VersionChecker) versionCheckLoop(
	ctx context.Context,
) {
	timer := time.NewTicker(VersionCheckInterval)
	defer timer.Stop()
	vc.performVersionCheck(ctx)
	for {
		select {
		case <-vc.shutdownChan:
			return
		case <-timer.C:
			vc.performVersionCheck(ctx)
		}
	}
}

func (vc *VersionChecker) performVersionCheck(
	ctx context.Context,
) {
	startTime := time.Now().UTC()
	defer func() {
		metrics.VersionCheckLatency.With(vc.metricsHandler).Record(time.Since(startTime))
	}()
	metadata, err := vc.clusterMetadataManager.GetCurrentClusterMetadata(ctx)
	if err != nil {
		metrics.VersionCheckFailedCount.With(vc.metricsHandler).Record(1)
		return
	}

	if !isUpdateNeeded(metadata) {
		return
	}

	req, err := vc.createVersionCheckRequest(metadata)
	if err != nil {
		metrics.VersionCheckFailedCount.With(vc.metricsHandler).Record(1)
		return
	}
	resp, err := vc.getVersionInfo(req)
	if err != nil {
		metrics.VersionCheckRequestFailedCount.With(vc.metricsHandler).Record(1)
		metrics.VersionCheckFailedCount.With(vc.metricsHandler).Record(1)
		return
	}
	err = vc.saveVersionInfo(ctx, resp)
	if err != nil {
		metrics.VersionCheckFailedCount.With(vc.metricsHandler).Record(1)
		return
	}
	metrics.VersionCheckSuccessCount.With(vc.metricsHandler).Record(1)
}

func isUpdateNeeded(metadata *persistence.GetClusterMetadataResponse) bool {
	return metadata.VersionInfo == nil || (metadata.VersionInfo.LastUpdateTime != nil &&
		metadata.VersionInfo.LastUpdateTime.AsTime().Before(time.Now().Add(-time.Hour)))
}

func (vc *VersionChecker) createVersionCheckRequest(metadata *persistence.GetClusterMetadataResponse) (*check.VersionCheckRequest, error) {
	return &check.VersionCheckRequest{
		Product:   headers.ClientNameServer,
		Version:   headers.ServerVersion,
		Arch:      runtime.GOARCH,
		OS:        runtime.GOOS,
		DB:        vc.clusterMetadataManager.GetName(),
		ClusterID: metadata.ClusterId,
		Timestamp: time.Now().UnixNano(),
		SDKInfo:   vc.sdkVersionRecorder.GetAndResetSDKInfo(),
	}, nil
}

func (vc *VersionChecker) getVersionInfo(req *check.VersionCheckRequest) (*check.VersionCheckResponse, error) {
	return check.NewCaller().Call(req)
}

func (vc *VersionChecker) saveVersionInfo(ctx context.Context, resp *check.VersionCheckResponse) error {
	metadata, err := vc.clusterMetadataManager.GetCurrentClusterMetadata(ctx)
	if err != nil {
		return err
	}
	// TODO(bergundy): Extract and save version info per SDK
	versionInfo, err := toVersionInfo(resp)
	if err != nil {
		return err
	}
	metadata.VersionInfo = versionInfo
	saved, err := vc.clusterMetadataManager.SaveClusterMetadata(ctx, &persistence.SaveClusterMetadataRequest{
		ClusterMetadata: metadata.ClusterMetadata, Version: metadata.Version})
	if err != nil {
		return err
	}
	if !saved {
		return serviceerror.NewUnavailable("version info update hasn't been applied")
	}
	return nil
}

func toVersionInfo(resp *check.VersionCheckResponse) (*versionpb.VersionInfo, error) {
	for _, product := range resp.Products {
		if product.Product == headers.ClientNameServer {
			return &versionpb.VersionInfo{
				Current:        convertReleaseInfo(product.Current),
				Recommended:    convertReleaseInfo(product.Recommended),
				Instructions:   product.Instructions,
				Alerts:         convertAlerts(product.Alerts),
				LastUpdateTime: timestamppb.New(time.Now().UTC()),
			}, nil
		}
	}
	return nil, serviceerror.NewNotFound("version info update was not found in response")
}

func convertAlerts(alerts []check.Alert) []*versionpb.Alert {
	var result []*versionpb.Alert
	for _, alert := range alerts {
		result = append(result, &versionpb.Alert{
			Message:  alert.Message,
			Severity: enumspb.Severity(alert.Severity),
		})
	}
	return result
}

func convertReleaseInfo(releaseInfo check.ReleaseInfo) *versionpb.ReleaseInfo {
	return &versionpb.ReleaseInfo{
		Version:     releaseInfo.Version,
		ReleaseTime: timestamp.UnixOrZeroTimePtr(releaseInfo.ReleaseTime),
		Notes:       releaseInfo.Notes,
	}
}
