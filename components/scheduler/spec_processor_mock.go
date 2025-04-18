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
// Source: spec_processor.go
//
// Generated by this command:
//
//	mockgen -copyright_file ../../LICENSE -package scheduler -source spec_processor.go -destination spec_processor_mock.go
//

// Package scheduler is a generated GoMock package.
package scheduler

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockSpecProcessor is a mock of SpecProcessor interface.
type MockSpecProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockSpecProcessorMockRecorder
	isgomock struct{}
}

// MockSpecProcessorMockRecorder is the mock recorder for MockSpecProcessor.
type MockSpecProcessorMockRecorder struct {
	mock *MockSpecProcessor
}

// NewMockSpecProcessor creates a new mock instance.
func NewMockSpecProcessor(ctrl *gomock.Controller) *MockSpecProcessor {
	mock := &MockSpecProcessor{ctrl: ctrl}
	mock.recorder = &MockSpecProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpecProcessor) EXPECT() *MockSpecProcessorMockRecorder {
	return m.recorder
}

// ProcessTimeRange mocks base method.
func (m *MockSpecProcessor) ProcessTimeRange(scheduler Scheduler, start, end time.Time, manual bool, limit *int) (*ProcessedTimeRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTimeRange", scheduler, start, end, manual, limit)
	ret0, _ := ret[0].(*ProcessedTimeRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessTimeRange indicates an expected call of ProcessTimeRange.
func (mr *MockSpecProcessorMockRecorder) ProcessTimeRange(scheduler, start, end, manual, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTimeRange", reflect.TypeOf((*MockSpecProcessor)(nil).ProcessTimeRange), scheduler, start, end, manual, limit)
}
