// Code generated by mockery v2.9.4. DO NOT EDIT.

package backend

import (
	context "context"

	core "github.com/cschleiden/go-dt/pkg/core"
	history "github.com/cschleiden/go-dt/pkg/history"

	mock "github.com/stretchr/testify/mock"

	task "github.com/cschleiden/go-dt/pkg/core/task"
)

// MockBackend is an autogenerated mock type for the Backend type
type MockBackend struct {
	mock.Mock
}

// CompleteActivityTask provides a mock function with given fields: ctx, instance, activityID, event
func (_m *MockBackend) CompleteActivityTask(ctx context.Context, instance core.WorkflowInstance, activityID string, event history.Event) error {
	ret := _m.Called(ctx, instance, activityID, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.WorkflowInstance, string, history.Event) error); ok {
		r0 = rf(ctx, instance, activityID, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteWorkflowTask provides a mock function with given fields: ctx, _a1, events, workflowEvents
func (_m *MockBackend) CompleteWorkflowTask(ctx context.Context, _a1 task.Workflow, events []history.Event, workflowEvents []core.WorkflowEvent) error {
	ret := _m.Called(ctx, _a1, events, workflowEvents)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, task.Workflow, []history.Event, []core.WorkflowEvent) error); ok {
		r0 = rf(ctx, _a1, events, workflowEvents)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateWorkflowInstance provides a mock function with given fields: ctx, event
func (_m *MockBackend) CreateWorkflowInstance(ctx context.Context, event core.WorkflowEvent) error {
	ret := _m.Called(ctx, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.WorkflowEvent) error); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtendActivityTask provides a mock function with given fields: ctx, activityID
func (_m *MockBackend) ExtendActivityTask(ctx context.Context, activityID string) error {
	ret := _m.Called(ctx, activityID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, activityID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtendWorkflowTask provides a mock function with given fields: ctx, instance
func (_m *MockBackend) ExtendWorkflowTask(ctx context.Context, instance core.WorkflowInstance) error {
	ret := _m.Called(ctx, instance)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.WorkflowInstance) error); ok {
		r0 = rf(ctx, instance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetActivityTask provides a mock function with given fields: ctx
func (_m *MockBackend) GetActivityTask(ctx context.Context) (*task.Activity, error) {
	ret := _m.Called(ctx)

	var r0 *task.Activity
	if rf, ok := ret.Get(0).(func(context.Context) *task.Activity); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*task.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowTask provides a mock function with given fields: ctx
func (_m *MockBackend) GetWorkflowTask(ctx context.Context) (*task.Workflow, error) {
	ret := _m.Called(ctx)

	var r0 *task.Workflow
	if rf, ok := ret.Get(0).(func(context.Context) *task.Workflow); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*task.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignalWorkflow provides a mock function with given fields: ctx, instance, event
func (_m *MockBackend) SignalWorkflow(ctx context.Context, instance core.WorkflowInstance, event history.Event) error {
	ret := _m.Called(ctx, instance, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.WorkflowInstance, history.Event) error); ok {
		r0 = rf(ctx, instance, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
