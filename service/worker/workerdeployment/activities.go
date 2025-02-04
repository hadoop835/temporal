// The MIT License
//
// Copyright (c) 2024 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2024 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package workerdeployment

import (
	"context"

	"go.temporal.io/sdk/activity"
	deploymentspb "go.temporal.io/server/api/deployment/v1"
	"go.temporal.io/server/common/namespace"
)

type (
	Activities struct {
		namespace        *namespace.Namespace
		deploymentClient Client
	}
)

func (a *Activities) SyncWorkerDeploymentVersion(ctx context.Context, args *deploymentspb.SyncVersionStateActivityArgs) (*deploymentspb.SyncVersionStateActivityResult, error) {
	identity := "worker-deployment workflow " + activity.GetInfo(ctx).WorkflowExecution.ID
	res, err := a.deploymentClient.SyncVersionWorkflowFromWorkerDeployment(
		ctx,
		a.namespace,
		args.DeploymentName,
		args.Version,
		args.UpdateArgs,
		identity,
		args.RequestId,
	)
	if err != nil {
		return nil, err
	}
	return &deploymentspb.SyncVersionStateActivityResult{
		VersionState: res.VersionState,
	}, nil
}

func (a *Activities) IsVersionMissingTaskQueues(ctx context.Context, args *deploymentspb.IsVersionMissingTaskQueuesArgs) (*deploymentspb.IsVersionMissingTaskQueuesResult, error) {
	res, err := a.deploymentClient.IsVersionMissingTaskQueues(
		ctx,
		a.namespace,
		args.PrevCurrentVersion,
		args.NewCurrentVersion,
	)
	if err != nil {
		return nil, err
	}
	return &deploymentspb.IsVersionMissingTaskQueuesResult{
		IsMissingTaskQueues: res,
	}, nil
}

func (a *Activities) DeleteWorkerDeploymentVersion(ctx context.Context, args *deploymentspb.DeleteVersionActivityArgs) error {
	identity := "worker-deployment workflow " + activity.GetInfo(ctx).WorkflowExecution.ID
	err := a.deploymentClient.DeleteVersionFromWorkerDeployment(
		ctx,
		a.namespace,
		args.DeploymentName,
		args.Version,
		identity,
		args.RequestId,
	)
	if err != nil {
		return err
	}
	return nil
}
