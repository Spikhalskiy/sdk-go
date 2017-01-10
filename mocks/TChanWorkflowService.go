package mocks

import minions "code.uber.internal/devexp/minions-client-go.git/.gen/go/minions"
import mock "github.com/stretchr/testify/mock"
import shared "code.uber.internal/devexp/minions-client-go.git/.gen/go/shared"
import thrift "github.com/uber/tchannel-go/thrift"

// TChanWorkflowService is an autogenerated mock type for the TChanWorkflowService type
type TChanWorkflowService struct {
	mock.Mock
}

// GetWorkflowExecutionHistory provides a mock function with given fields: ctx, getRequest
func (_m *TChanWorkflowService) GetWorkflowExecutionHistory(ctx thrift.Context, getRequest *shared.GetWorkflowExecutionHistoryRequest) (*shared.GetWorkflowExecutionHistoryResponse, error) {
	ret := _m.Called(ctx, getRequest)

	var r0 *shared.GetWorkflowExecutionHistoryResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.GetWorkflowExecutionHistoryRequest) *shared.GetWorkflowExecutionHistoryResponse); ok {
		r0 = rf(ctx, getRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.GetWorkflowExecutionHistoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.GetWorkflowExecutionHistoryRequest) error); ok {
		r1 = rf(ctx, getRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PollForActivityTask provides a mock function with given fields: ctx, pollRequest
func (_m *TChanWorkflowService) PollForActivityTask(ctx thrift.Context, pollRequest *shared.PollForActivityTaskRequest) (*shared.PollForActivityTaskResponse, error) {
	ret := _m.Called(ctx, pollRequest)

	var r0 *shared.PollForActivityTaskResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.PollForActivityTaskRequest) *shared.PollForActivityTaskResponse); ok {
		r0 = rf(ctx, pollRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.PollForActivityTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.PollForActivityTaskRequest) error); ok {
		r1 = rf(ctx, pollRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PollForDecisionTask provides a mock function with given fields: ctx, pollRequest
func (_m *TChanWorkflowService) PollForDecisionTask(ctx thrift.Context, pollRequest *shared.PollForDecisionTaskRequest) (*shared.PollForDecisionTaskResponse, error) {
	ret := _m.Called(ctx, pollRequest)

	var r0 *shared.PollForDecisionTaskResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.PollForDecisionTaskRequest) *shared.PollForDecisionTaskResponse); ok {
		r0 = rf(ctx, pollRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.PollForDecisionTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.PollForDecisionTaskRequest) error); ok {
		r1 = rf(ctx, pollRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecordActivityTaskHeartbeat provides a mock function with given fields: ctx, heartbeatRequest
func (_m *TChanWorkflowService) RecordActivityTaskHeartbeat(ctx thrift.Context, heartbeatRequest *shared.RecordActivityTaskHeartbeatRequest) (*shared.RecordActivityTaskHeartbeatResponse, error) {
	ret := _m.Called(ctx, heartbeatRequest)

	var r0 *shared.RecordActivityTaskHeartbeatResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.RecordActivityTaskHeartbeatRequest) *shared.RecordActivityTaskHeartbeatResponse); ok {
		r0 = rf(ctx, heartbeatRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.RecordActivityTaskHeartbeatResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.RecordActivityTaskHeartbeatRequest) error); ok {
		r1 = rf(ctx, heartbeatRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RespondActivityTaskCompleted provides a mock function with given fields: ctx, completeRequest
func (_m *TChanWorkflowService) RespondActivityTaskCompleted(ctx thrift.Context, completeRequest *shared.RespondActivityTaskCompletedRequest) error {
	ret := _m.Called(ctx, completeRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.RespondActivityTaskCompletedRequest) error); ok {
		r0 = rf(ctx, completeRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondActivityTaskFailed provides a mock function with given fields: ctx, failRequest
func (_m *TChanWorkflowService) RespondActivityTaskFailed(ctx thrift.Context, failRequest *shared.RespondActivityTaskFailedRequest) error {
	ret := _m.Called(ctx, failRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.RespondActivityTaskFailedRequest) error); ok {
		r0 = rf(ctx, failRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondDecisionTaskCompleted provides a mock function with given fields: ctx, completeRequest
func (_m *TChanWorkflowService) RespondDecisionTaskCompleted(ctx thrift.Context, completeRequest *shared.RespondDecisionTaskCompletedRequest) error {
	ret := _m.Called(ctx, completeRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.RespondDecisionTaskCompletedRequest) error); ok {
		r0 = rf(ctx, completeRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartWorkflowExecution provides a mock function with given fields: ctx, startRequest
func (_m *TChanWorkflowService) StartWorkflowExecution(ctx thrift.Context, startRequest *shared.StartWorkflowExecutionRequest) (*shared.StartWorkflowExecutionResponse, error) {
	ret := _m.Called(ctx, startRequest)

	var r0 *shared.StartWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.StartWorkflowExecutionRequest) *shared.StartWorkflowExecutionResponse); ok {
		r0 = rf(ctx, startRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.StartWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.StartWorkflowExecutionRequest) error); ok {
		r1 = rf(ctx, startRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ minions.TChanWorkflowService = (*TChanWorkflowService)(nil)