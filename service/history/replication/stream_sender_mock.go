// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
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

// Code generated by MockGen. DO NOT EDIT.
// Source: stream_sender.go
//
// Generated by this command:
//
//	mockgen -copyright_file ../../../LICENSE -package replication -source stream_sender.go -destination stream_sender_mock.go
//

// Package replication is a generated GoMock package.
package replication

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStreamSender is a mock of StreamSender interface.
type MockStreamSender struct {
	ctrl     *gomock.Controller
	recorder *MockStreamSenderMockRecorder
	isgomock struct{}
}

// MockStreamSenderMockRecorder is the mock recorder for MockStreamSender.
type MockStreamSenderMockRecorder struct {
	mock *MockStreamSender
}

// NewMockStreamSender creates a new mock instance.
func NewMockStreamSender(ctrl *gomock.Controller) *MockStreamSender {
	mock := &MockStreamSender{ctrl: ctrl}
	mock.recorder = &MockStreamSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamSender) EXPECT() *MockStreamSenderMockRecorder {
	return m.recorder
}

// IsValid mocks base method.
func (m *MockStreamSender) IsValid() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValid indicates an expected call of IsValid.
func (mr *MockStreamSenderMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockStreamSender)(nil).IsValid))
}

// Key mocks base method.
func (m *MockStreamSender) Key() ClusterShardKeyPair {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].(ClusterShardKeyPair)
	return ret0
}

// Key indicates an expected call of Key.
func (mr *MockStreamSenderMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockStreamSender)(nil).Key))
}

// Stop mocks base method.
func (m *MockStreamSender) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockStreamSenderMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockStreamSender)(nil).Stop))
}
