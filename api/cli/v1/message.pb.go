// The MIT License
//
// Copyright (c) 2020 Temporal Technologies, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/server/api/cli/v1/message.proto

package cli

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	v11 "go.temporal.io/api/common/v1"
	v12 "go.temporal.io/api/enums/v1"
	v1 "go.temporal.io/api/workflow/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DescribeWorkflowExecutionResponse struct {
	state                 protoimpl.MessageState          `protogen:"open.v1"`
	ExecutionConfig       *v1.WorkflowExecutionConfig     `protobuf:"bytes,1,opt,name=execution_config,json=executionConfig,proto3" json:"execution_config,omitempty"`
	WorkflowExecutionInfo *WorkflowExecutionInfo          `protobuf:"bytes,2,opt,name=workflow_execution_info,json=workflowExecutionInfo,proto3" json:"workflow_execution_info,omitempty"`
	PendingActivities     []*PendingActivityInfo          `protobuf:"bytes,3,rep,name=pending_activities,json=pendingActivities,proto3" json:"pending_activities,omitempty"`
	PendingChildren       []*v1.PendingChildExecutionInfo `protobuf:"bytes,4,rep,name=pending_children,json=pendingChildren,proto3" json:"pending_children,omitempty"`
	PendingWorkflowTask   *v1.PendingWorkflowTaskInfo     `protobuf:"bytes,5,opt,name=pending_workflow_task,json=pendingWorkflowTask,proto3" json:"pending_workflow_task,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *DescribeWorkflowExecutionResponse) Reset() {
	*x = DescribeWorkflowExecutionResponse{}
	mi := &file_temporal_server_api_cli_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeWorkflowExecutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeWorkflowExecutionResponse) ProtoMessage() {}

func (x *DescribeWorkflowExecutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_cli_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeWorkflowExecutionResponse.ProtoReflect.Descriptor instead.
func (*DescribeWorkflowExecutionResponse) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_cli_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *DescribeWorkflowExecutionResponse) GetExecutionConfig() *v1.WorkflowExecutionConfig {
	if x != nil {
		return x.ExecutionConfig
	}
	return nil
}

func (x *DescribeWorkflowExecutionResponse) GetWorkflowExecutionInfo() *WorkflowExecutionInfo {
	if x != nil {
		return x.WorkflowExecutionInfo
	}
	return nil
}

func (x *DescribeWorkflowExecutionResponse) GetPendingActivities() []*PendingActivityInfo {
	if x != nil {
		return x.PendingActivities
	}
	return nil
}

func (x *DescribeWorkflowExecutionResponse) GetPendingChildren() []*v1.PendingChildExecutionInfo {
	if x != nil {
		return x.PendingChildren
	}
	return nil
}

func (x *DescribeWorkflowExecutionResponse) GetPendingWorkflowTask() *v1.PendingWorkflowTaskInfo {
	if x != nil {
		return x.PendingWorkflowTask
	}
	return nil
}

type WorkflowExecutionInfo struct {
	state                        protoimpl.MessageState      `protogen:"open.v1"`
	Execution                    *v11.WorkflowExecution      `protobuf:"bytes,1,opt,name=execution,proto3" json:"execution,omitempty"`
	Type                         *v11.WorkflowType           `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	StartTime                    *timestamppb.Timestamp      `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	CloseTime                    *timestamppb.Timestamp      `protobuf:"bytes,4,opt,name=close_time,json=closeTime,proto3" json:"close_time,omitempty"`
	Status                       v12.WorkflowExecutionStatus `protobuf:"varint,5,opt,name=status,proto3,enum=temporal.api.enums.v1.WorkflowExecutionStatus" json:"status,omitempty"`
	HistoryLength                int64                       `protobuf:"varint,6,opt,name=history_length,json=historyLength,proto3" json:"history_length,omitempty"`
	ParentNamespaceId            string                      `protobuf:"bytes,7,opt,name=parent_namespace_id,json=parentNamespaceId,proto3" json:"parent_namespace_id,omitempty"`
	ParentExecution              *v11.WorkflowExecution      `protobuf:"bytes,8,opt,name=parent_execution,json=parentExecution,proto3" json:"parent_execution,omitempty"`
	ExecutionTime                *timestamppb.Timestamp      `protobuf:"bytes,9,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	Memo                         *v11.Memo                   `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo,omitempty"`
	SearchAttributes             *SearchAttributes           `protobuf:"bytes,11,opt,name=search_attributes,json=searchAttributes,proto3" json:"search_attributes,omitempty"`
	AutoResetPoints              *v1.ResetPoints             `protobuf:"bytes,12,opt,name=auto_reset_points,json=autoResetPoints,proto3" json:"auto_reset_points,omitempty"`
	StateTransitionCount         int64                       `protobuf:"varint,13,opt,name=state_transition_count,json=stateTransitionCount,proto3" json:"state_transition_count,omitempty"`
	HistorySizeBytes             int64                       `protobuf:"varint,14,opt,name=history_size_bytes,json=historySizeBytes,proto3" json:"history_size_bytes,omitempty"`
	MostRecentWorkerVersionStamp *v11.WorkerVersionStamp     `protobuf:"bytes,15,opt,name=most_recent_worker_version_stamp,json=mostRecentWorkerVersionStamp,proto3" json:"most_recent_worker_version_stamp,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *WorkflowExecutionInfo) Reset() {
	*x = WorkflowExecutionInfo{}
	mi := &file_temporal_server_api_cli_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowExecutionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowExecutionInfo) ProtoMessage() {}

func (x *WorkflowExecutionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_cli_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowExecutionInfo.ProtoReflect.Descriptor instead.
func (*WorkflowExecutionInfo) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_cli_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowExecutionInfo) GetExecution() *v11.WorkflowExecution {
	if x != nil {
		return x.Execution
	}
	return nil
}

func (x *WorkflowExecutionInfo) GetType() *v11.WorkflowType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *WorkflowExecutionInfo) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *WorkflowExecutionInfo) GetCloseTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CloseTime
	}
	return nil
}

func (x *WorkflowExecutionInfo) GetStatus() v12.WorkflowExecutionStatus {
	if x != nil {
		return x.Status
	}
	return v12.WorkflowExecutionStatus(0)
}

func (x *WorkflowExecutionInfo) GetHistoryLength() int64 {
	if x != nil {
		return x.HistoryLength
	}
	return 0
}

func (x *WorkflowExecutionInfo) GetParentNamespaceId() string {
	if x != nil {
		return x.ParentNamespaceId
	}
	return ""
}

func (x *WorkflowExecutionInfo) GetParentExecution() *v11.WorkflowExecution {
	if x != nil {
		return x.ParentExecution
	}
	return nil
}

func (x *WorkflowExecutionInfo) GetExecutionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExecutionTime
	}
	return nil
}

func (x *WorkflowExecutionInfo) GetMemo() *v11.Memo {
	if x != nil {
		return x.Memo
	}
	return nil
}

func (x *WorkflowExecutionInfo) GetSearchAttributes() *SearchAttributes {
	if x != nil {
		return x.SearchAttributes
	}
	return nil
}

func (x *WorkflowExecutionInfo) GetAutoResetPoints() *v1.ResetPoints {
	if x != nil {
		return x.AutoResetPoints
	}
	return nil
}

func (x *WorkflowExecutionInfo) GetStateTransitionCount() int64 {
	if x != nil {
		return x.StateTransitionCount
	}
	return 0
}

func (x *WorkflowExecutionInfo) GetHistorySizeBytes() int64 {
	if x != nil {
		return x.HistorySizeBytes
	}
	return 0
}

func (x *WorkflowExecutionInfo) GetMostRecentWorkerVersionStamp() *v11.WorkerVersionStamp {
	if x != nil {
		return x.MostRecentWorkerVersionStamp
	}
	return nil
}

type PendingActivityInfo struct {
	state              protoimpl.MessageState   `protogen:"open.v1"`
	ActivityId         string                   `protobuf:"bytes,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	ActivityType       *v11.ActivityType        `protobuf:"bytes,2,opt,name=activity_type,json=activityType,proto3" json:"activity_type,omitempty"`
	State              v12.PendingActivityState `protobuf:"varint,3,opt,name=state,proto3,enum=temporal.api.enums.v1.PendingActivityState" json:"state,omitempty"`
	HeartbeatDetails   string                   `protobuf:"bytes,4,opt,name=heartbeat_details,json=heartbeatDetails,proto3" json:"heartbeat_details,omitempty"`
	LastHeartbeatTime  *timestamppb.Timestamp   `protobuf:"bytes,5,opt,name=last_heartbeat_time,json=lastHeartbeatTime,proto3" json:"last_heartbeat_time,omitempty"`
	LastStartedTime    *timestamppb.Timestamp   `protobuf:"bytes,6,opt,name=last_started_time,json=lastStartedTime,proto3" json:"last_started_time,omitempty"`
	Attempt            int32                    `protobuf:"varint,7,opt,name=attempt,proto3" json:"attempt,omitempty"`
	MaximumAttempts    int32                    `protobuf:"varint,8,opt,name=maximum_attempts,json=maximumAttempts,proto3" json:"maximum_attempts,omitempty"`
	ScheduledTime      *timestamppb.Timestamp   `protobuf:"bytes,9,opt,name=scheduled_time,json=scheduledTime,proto3" json:"scheduled_time,omitempty"`
	ExpirationTime     *timestamppb.Timestamp   `protobuf:"bytes,10,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	LastFailure        *Failure                 `protobuf:"bytes,11,opt,name=last_failure,json=lastFailure,proto3" json:"last_failure,omitempty"`
	LastWorkerIdentity string                   `protobuf:"bytes,12,opt,name=last_worker_identity,json=lastWorkerIdentity,proto3" json:"last_worker_identity,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *PendingActivityInfo) Reset() {
	*x = PendingActivityInfo{}
	mi := &file_temporal_server_api_cli_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PendingActivityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingActivityInfo) ProtoMessage() {}

func (x *PendingActivityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_cli_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingActivityInfo.ProtoReflect.Descriptor instead.
func (*PendingActivityInfo) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_cli_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *PendingActivityInfo) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

func (x *PendingActivityInfo) GetActivityType() *v11.ActivityType {
	if x != nil {
		return x.ActivityType
	}
	return nil
}

func (x *PendingActivityInfo) GetState() v12.PendingActivityState {
	if x != nil {
		return x.State
	}
	return v12.PendingActivityState(0)
}

func (x *PendingActivityInfo) GetHeartbeatDetails() string {
	if x != nil {
		return x.HeartbeatDetails
	}
	return ""
}

func (x *PendingActivityInfo) GetLastHeartbeatTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastHeartbeatTime
	}
	return nil
}

func (x *PendingActivityInfo) GetLastStartedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastStartedTime
	}
	return nil
}

func (x *PendingActivityInfo) GetAttempt() int32 {
	if x != nil {
		return x.Attempt
	}
	return 0
}

func (x *PendingActivityInfo) GetMaximumAttempts() int32 {
	if x != nil {
		return x.MaximumAttempts
	}
	return 0
}

func (x *PendingActivityInfo) GetScheduledTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledTime
	}
	return nil
}

func (x *PendingActivityInfo) GetExpirationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *PendingActivityInfo) GetLastFailure() *Failure {
	if x != nil {
		return x.LastFailure
	}
	return nil
}

func (x *PendingActivityInfo) GetLastWorkerIdentity() string {
	if x != nil {
		return x.LastWorkerIdentity
	}
	return ""
}

type SearchAttributes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IndexedFields map[string]string      `protobuf:"bytes,1,rep,name=indexed_fields,json=indexedFields,proto3" json:"indexed_fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchAttributes) Reset() {
	*x = SearchAttributes{}
	mi := &file_temporal_server_api_cli_v1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchAttributes) ProtoMessage() {}

func (x *SearchAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_cli_v1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchAttributes.ProtoReflect.Descriptor instead.
func (*SearchAttributes) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_cli_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *SearchAttributes) GetIndexedFields() map[string]string {
	if x != nil {
		return x.IndexedFields
	}
	return nil
}

type Failure struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Source        string                 `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	StackTrace    string                 `protobuf:"bytes,3,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	Cause         *Failure               `protobuf:"bytes,4,opt,name=cause,proto3" json:"cause,omitempty"`
	FailureType   string                 `protobuf:"bytes,5,opt,name=failure_type,json=failureType,proto3" json:"failure_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Failure) Reset() {
	*x = Failure{}
	mi := &file_temporal_server_api_cli_v1_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Failure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Failure) ProtoMessage() {}

func (x *Failure) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_cli_v1_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Failure.ProtoReflect.Descriptor instead.
func (*Failure) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_cli_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *Failure) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Failure) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Failure) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

func (x *Failure) GetCause() *Failure {
	if x != nil {
		return x.Cause
	}
	return nil
}

func (x *Failure) GetFailureType() string {
	if x != nil {
		return x.FailureType
	}
	return ""
}

type AddSearchAttributesResponse struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	IndexName                string                 `protobuf:"bytes,1,opt,name=index_name,json=indexName,proto3" json:"index_name,omitempty"`
	CustomSearchAttributes   map[string]string      `protobuf:"bytes,2,rep,name=custom_search_attributes,json=customSearchAttributes,proto3" json:"custom_search_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SystemSearchAttributes   map[string]string      `protobuf:"bytes,3,rep,name=system_search_attributes,json=systemSearchAttributes,proto3" json:"system_search_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Mapping                  map[string]string      `protobuf:"bytes,4,rep,name=mapping,proto3" json:"mapping,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	AddWorkflowExecutionInfo *WorkflowExecutionInfo `protobuf:"bytes,5,opt,name=add_workflow_execution_info,json=addWorkflowExecutionInfo,proto3" json:"add_workflow_execution_info,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *AddSearchAttributesResponse) Reset() {
	*x = AddSearchAttributesResponse{}
	mi := &file_temporal_server_api_cli_v1_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddSearchAttributesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSearchAttributesResponse) ProtoMessage() {}

func (x *AddSearchAttributesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_cli_v1_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSearchAttributesResponse.ProtoReflect.Descriptor instead.
func (*AddSearchAttributesResponse) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_cli_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *AddSearchAttributesResponse) GetIndexName() string {
	if x != nil {
		return x.IndexName
	}
	return ""
}

func (x *AddSearchAttributesResponse) GetCustomSearchAttributes() map[string]string {
	if x != nil {
		return x.CustomSearchAttributes
	}
	return nil
}

func (x *AddSearchAttributesResponse) GetSystemSearchAttributes() map[string]string {
	if x != nil {
		return x.SystemSearchAttributes
	}
	return nil
}

func (x *AddSearchAttributesResponse) GetMapping() map[string]string {
	if x != nil {
		return x.Mapping
	}
	return nil
}

func (x *AddSearchAttributesResponse) GetAddWorkflowExecutionInfo() *WorkflowExecutionInfo {
	if x != nil {
		return x.AddWorkflowExecutionInfo
	}
	return nil
}

var File_temporal_server_api_cli_v1_message_proto protoreflect.FileDescriptor

const file_temporal_server_api_cli_v1_message_proto_rawDesc = "" +
	"\n" +
	"(temporal/server/api/cli/v1/message.proto\x12\x1atemporal.server.api.cli.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a$temporal/api/common/v1/message.proto\x1a$temporal/api/enums/v1/workflow.proto\x1a&temporal/api/workflow/v1/message.proto\"\x93\x04\n" +
	"!DescribeWorkflowExecutionResponse\x12\\\n" +
	"\x10execution_config\x18\x01 \x01(\v21.temporal.api.workflow.v1.WorkflowExecutionConfigR\x0fexecutionConfig\x12i\n" +
	"\x17workflow_execution_info\x18\x02 \x01(\v21.temporal.server.api.cli.v1.WorkflowExecutionInfoR\x15workflowExecutionInfo\x12^\n" +
	"\x12pending_activities\x18\x03 \x03(\v2/.temporal.server.api.cli.v1.PendingActivityInfoR\x11pendingActivities\x12^\n" +
	"\x10pending_children\x18\x04 \x03(\v23.temporal.api.workflow.v1.PendingChildExecutionInfoR\x0fpendingChildren\x12e\n" +
	"\x15pending_workflow_task\x18\x05 \x01(\v21.temporal.api.workflow.v1.PendingWorkflowTaskInfoR\x13pendingWorkflowTask\"\x80\b\n" +
	"\x15WorkflowExecutionInfo\x12G\n" +
	"\texecution\x18\x01 \x01(\v2).temporal.api.common.v1.WorkflowExecutionR\texecution\x128\n" +
	"\x04type\x18\x02 \x01(\v2$.temporal.api.common.v1.WorkflowTypeR\x04type\x129\n" +
	"\n" +
	"start_time\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tstartTime\x129\n" +
	"\n" +
	"close_time\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tcloseTime\x12F\n" +
	"\x06status\x18\x05 \x01(\x0e2..temporal.api.enums.v1.WorkflowExecutionStatusR\x06status\x12%\n" +
	"\x0ehistory_length\x18\x06 \x01(\x03R\rhistoryLength\x12.\n" +
	"\x13parent_namespace_id\x18\a \x01(\tR\x11parentNamespaceId\x12T\n" +
	"\x10parent_execution\x18\b \x01(\v2).temporal.api.common.v1.WorkflowExecutionR\x0fparentExecution\x12A\n" +
	"\x0eexecution_time\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\rexecutionTime\x120\n" +
	"\x04memo\x18\n" +
	" \x01(\v2\x1c.temporal.api.common.v1.MemoR\x04memo\x12Y\n" +
	"\x11search_attributes\x18\v \x01(\v2,.temporal.server.api.cli.v1.SearchAttributesR\x10searchAttributes\x12Q\n" +
	"\x11auto_reset_points\x18\f \x01(\v2%.temporal.api.workflow.v1.ResetPointsR\x0fautoResetPoints\x124\n" +
	"\x16state_transition_count\x18\r \x01(\x03R\x14stateTransitionCount\x12,\n" +
	"\x12history_size_bytes\x18\x0e \x01(\x03R\x10historySizeBytes\x12r\n" +
	" most_recent_worker_version_stamp\x18\x0f \x01(\v2*.temporal.api.common.v1.WorkerVersionStampR\x1cmostRecentWorkerVersionStamp\"\xcc\x05\n" +
	"\x13PendingActivityInfo\x12\x1f\n" +
	"\vactivity_id\x18\x01 \x01(\tR\n" +
	"activityId\x12I\n" +
	"\ractivity_type\x18\x02 \x01(\v2$.temporal.api.common.v1.ActivityTypeR\factivityType\x12A\n" +
	"\x05state\x18\x03 \x01(\x0e2+.temporal.api.enums.v1.PendingActivityStateR\x05state\x12+\n" +
	"\x11heartbeat_details\x18\x04 \x01(\tR\x10heartbeatDetails\x12J\n" +
	"\x13last_heartbeat_time\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\x11lastHeartbeatTime\x12F\n" +
	"\x11last_started_time\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\x0flastStartedTime\x12\x18\n" +
	"\aattempt\x18\a \x01(\x05R\aattempt\x12)\n" +
	"\x10maximum_attempts\x18\b \x01(\x05R\x0fmaximumAttempts\x12A\n" +
	"\x0escheduled_time\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\rscheduledTime\x12C\n" +
	"\x0fexpiration_time\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampR\x0eexpirationTime\x12F\n" +
	"\flast_failure\x18\v \x01(\v2#.temporal.server.api.cli.v1.FailureR\vlastFailure\x120\n" +
	"\x14last_worker_identity\x18\f \x01(\tR\x12lastWorkerIdentity\"\xbc\x01\n" +
	"\x10SearchAttributes\x12f\n" +
	"\x0eindexed_fields\x18\x01 \x03(\v2?.temporal.server.api.cli.v1.SearchAttributes.IndexedFieldsEntryR\rindexedFields\x1a@\n" +
	"\x12IndexedFieldsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xba\x01\n" +
	"\aFailure\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12\x16\n" +
	"\x06source\x18\x02 \x01(\tR\x06source\x12\x1f\n" +
	"\vstack_trace\x18\x03 \x01(\tR\n" +
	"stackTrace\x129\n" +
	"\x05cause\x18\x04 \x01(\v2#.temporal.server.api.cli.v1.FailureR\x05cause\x12!\n" +
	"\ffailure_type\x18\x05 \x01(\tR\vfailureType\"\x80\x06\n" +
	"\x1bAddSearchAttributesResponse\x12\x1d\n" +
	"\n" +
	"index_name\x18\x01 \x01(\tR\tindexName\x12\x8d\x01\n" +
	"\x18custom_search_attributes\x18\x02 \x03(\v2S.temporal.server.api.cli.v1.AddSearchAttributesResponse.CustomSearchAttributesEntryR\x16customSearchAttributes\x12\x8d\x01\n" +
	"\x18system_search_attributes\x18\x03 \x03(\v2S.temporal.server.api.cli.v1.AddSearchAttributesResponse.SystemSearchAttributesEntryR\x16systemSearchAttributes\x12^\n" +
	"\amapping\x18\x04 \x03(\v2D.temporal.server.api.cli.v1.AddSearchAttributesResponse.MappingEntryR\amapping\x12p\n" +
	"\x1badd_workflow_execution_info\x18\x05 \x01(\v21.temporal.server.api.cli.v1.WorkflowExecutionInfoR\x18addWorkflowExecutionInfo\x1aI\n" +
	"\x1bCustomSearchAttributesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1aI\n" +
	"\x1bSystemSearchAttributesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1a:\n" +
	"\fMappingEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B&Z$go.temporal.io/server/api/cli/v1;clib\x06proto3"

var (
	file_temporal_server_api_cli_v1_message_proto_rawDescOnce sync.Once
	file_temporal_server_api_cli_v1_message_proto_rawDescData []byte
)

func file_temporal_server_api_cli_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_server_api_cli_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_cli_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_server_api_cli_v1_message_proto_rawDesc), len(file_temporal_server_api_cli_v1_message_proto_rawDesc)))
	})
	return file_temporal_server_api_cli_v1_message_proto_rawDescData
}

var file_temporal_server_api_cli_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_temporal_server_api_cli_v1_message_proto_goTypes = []any{
	(*DescribeWorkflowExecutionResponse)(nil), // 0: temporal.server.api.cli.v1.DescribeWorkflowExecutionResponse
	(*WorkflowExecutionInfo)(nil),             // 1: temporal.server.api.cli.v1.WorkflowExecutionInfo
	(*PendingActivityInfo)(nil),               // 2: temporal.server.api.cli.v1.PendingActivityInfo
	(*SearchAttributes)(nil),                  // 3: temporal.server.api.cli.v1.SearchAttributes
	(*Failure)(nil),                           // 4: temporal.server.api.cli.v1.Failure
	(*AddSearchAttributesResponse)(nil),       // 5: temporal.server.api.cli.v1.AddSearchAttributesResponse
	nil,                                       // 6: temporal.server.api.cli.v1.SearchAttributes.IndexedFieldsEntry
	nil,                                       // 7: temporal.server.api.cli.v1.AddSearchAttributesResponse.CustomSearchAttributesEntry
	nil,                                       // 8: temporal.server.api.cli.v1.AddSearchAttributesResponse.SystemSearchAttributesEntry
	nil,                                       // 9: temporal.server.api.cli.v1.AddSearchAttributesResponse.MappingEntry
	(*v1.WorkflowExecutionConfig)(nil),        // 10: temporal.api.workflow.v1.WorkflowExecutionConfig
	(*v1.PendingChildExecutionInfo)(nil),      // 11: temporal.api.workflow.v1.PendingChildExecutionInfo
	(*v1.PendingWorkflowTaskInfo)(nil),        // 12: temporal.api.workflow.v1.PendingWorkflowTaskInfo
	(*v11.WorkflowExecution)(nil),             // 13: temporal.api.common.v1.WorkflowExecution
	(*v11.WorkflowType)(nil),                  // 14: temporal.api.common.v1.WorkflowType
	(*timestamppb.Timestamp)(nil),             // 15: google.protobuf.Timestamp
	(v12.WorkflowExecutionStatus)(0),          // 16: temporal.api.enums.v1.WorkflowExecutionStatus
	(*v11.Memo)(nil),                          // 17: temporal.api.common.v1.Memo
	(*v1.ResetPoints)(nil),                    // 18: temporal.api.workflow.v1.ResetPoints
	(*v11.WorkerVersionStamp)(nil),            // 19: temporal.api.common.v1.WorkerVersionStamp
	(*v11.ActivityType)(nil),                  // 20: temporal.api.common.v1.ActivityType
	(v12.PendingActivityState)(0),             // 21: temporal.api.enums.v1.PendingActivityState
}
var file_temporal_server_api_cli_v1_message_proto_depIdxs = []int32{
	10, // 0: temporal.server.api.cli.v1.DescribeWorkflowExecutionResponse.execution_config:type_name -> temporal.api.workflow.v1.WorkflowExecutionConfig
	1,  // 1: temporal.server.api.cli.v1.DescribeWorkflowExecutionResponse.workflow_execution_info:type_name -> temporal.server.api.cli.v1.WorkflowExecutionInfo
	2,  // 2: temporal.server.api.cli.v1.DescribeWorkflowExecutionResponse.pending_activities:type_name -> temporal.server.api.cli.v1.PendingActivityInfo
	11, // 3: temporal.server.api.cli.v1.DescribeWorkflowExecutionResponse.pending_children:type_name -> temporal.api.workflow.v1.PendingChildExecutionInfo
	12, // 4: temporal.server.api.cli.v1.DescribeWorkflowExecutionResponse.pending_workflow_task:type_name -> temporal.api.workflow.v1.PendingWorkflowTaskInfo
	13, // 5: temporal.server.api.cli.v1.WorkflowExecutionInfo.execution:type_name -> temporal.api.common.v1.WorkflowExecution
	14, // 6: temporal.server.api.cli.v1.WorkflowExecutionInfo.type:type_name -> temporal.api.common.v1.WorkflowType
	15, // 7: temporal.server.api.cli.v1.WorkflowExecutionInfo.start_time:type_name -> google.protobuf.Timestamp
	15, // 8: temporal.server.api.cli.v1.WorkflowExecutionInfo.close_time:type_name -> google.protobuf.Timestamp
	16, // 9: temporal.server.api.cli.v1.WorkflowExecutionInfo.status:type_name -> temporal.api.enums.v1.WorkflowExecutionStatus
	13, // 10: temporal.server.api.cli.v1.WorkflowExecutionInfo.parent_execution:type_name -> temporal.api.common.v1.WorkflowExecution
	15, // 11: temporal.server.api.cli.v1.WorkflowExecutionInfo.execution_time:type_name -> google.protobuf.Timestamp
	17, // 12: temporal.server.api.cli.v1.WorkflowExecutionInfo.memo:type_name -> temporal.api.common.v1.Memo
	3,  // 13: temporal.server.api.cli.v1.WorkflowExecutionInfo.search_attributes:type_name -> temporal.server.api.cli.v1.SearchAttributes
	18, // 14: temporal.server.api.cli.v1.WorkflowExecutionInfo.auto_reset_points:type_name -> temporal.api.workflow.v1.ResetPoints
	19, // 15: temporal.server.api.cli.v1.WorkflowExecutionInfo.most_recent_worker_version_stamp:type_name -> temporal.api.common.v1.WorkerVersionStamp
	20, // 16: temporal.server.api.cli.v1.PendingActivityInfo.activity_type:type_name -> temporal.api.common.v1.ActivityType
	21, // 17: temporal.server.api.cli.v1.PendingActivityInfo.state:type_name -> temporal.api.enums.v1.PendingActivityState
	15, // 18: temporal.server.api.cli.v1.PendingActivityInfo.last_heartbeat_time:type_name -> google.protobuf.Timestamp
	15, // 19: temporal.server.api.cli.v1.PendingActivityInfo.last_started_time:type_name -> google.protobuf.Timestamp
	15, // 20: temporal.server.api.cli.v1.PendingActivityInfo.scheduled_time:type_name -> google.protobuf.Timestamp
	15, // 21: temporal.server.api.cli.v1.PendingActivityInfo.expiration_time:type_name -> google.protobuf.Timestamp
	4,  // 22: temporal.server.api.cli.v1.PendingActivityInfo.last_failure:type_name -> temporal.server.api.cli.v1.Failure
	6,  // 23: temporal.server.api.cli.v1.SearchAttributes.indexed_fields:type_name -> temporal.server.api.cli.v1.SearchAttributes.IndexedFieldsEntry
	4,  // 24: temporal.server.api.cli.v1.Failure.cause:type_name -> temporal.server.api.cli.v1.Failure
	7,  // 25: temporal.server.api.cli.v1.AddSearchAttributesResponse.custom_search_attributes:type_name -> temporal.server.api.cli.v1.AddSearchAttributesResponse.CustomSearchAttributesEntry
	8,  // 26: temporal.server.api.cli.v1.AddSearchAttributesResponse.system_search_attributes:type_name -> temporal.server.api.cli.v1.AddSearchAttributesResponse.SystemSearchAttributesEntry
	9,  // 27: temporal.server.api.cli.v1.AddSearchAttributesResponse.mapping:type_name -> temporal.server.api.cli.v1.AddSearchAttributesResponse.MappingEntry
	1,  // 28: temporal.server.api.cli.v1.AddSearchAttributesResponse.add_workflow_execution_info:type_name -> temporal.server.api.cli.v1.WorkflowExecutionInfo
	29, // [29:29] is the sub-list for method output_type
	29, // [29:29] is the sub-list for method input_type
	29, // [29:29] is the sub-list for extension type_name
	29, // [29:29] is the sub-list for extension extendee
	0,  // [0:29] is the sub-list for field type_name
}

func init() { file_temporal_server_api_cli_v1_message_proto_init() }
func file_temporal_server_api_cli_v1_message_proto_init() {
	if File_temporal_server_api_cli_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_server_api_cli_v1_message_proto_rawDesc), len(file_temporal_server_api_cli_v1_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_cli_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_cli_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_server_api_cli_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_server_api_cli_v1_message_proto = out.File
	file_temporal_server_api_cli_v1_message_proto_goTypes = nil
	file_temporal_server_api_cli_v1_message_proto_depIdxs = nil
}
