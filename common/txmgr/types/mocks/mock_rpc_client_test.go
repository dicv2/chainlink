// Code generated by mockery v2.43.2. DO NOT EDIT.

package types

import (
	context "context"

	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// TxStrategy is an autogenerated mock type for the TxStrategy type
type TxStrategy struct {
	mock.Mock
}

type TxStrategy_Expecter struct {
	mock *mock.Mock
}

func (_m *TxStrategy) EXPECT() *TxStrategy_Expecter {
	return &TxStrategy_Expecter{mock: &_m.Mock}
}

// PruneQueue provides a mock function with given fields: ctx, pruneService
func (_m *TxStrategy) PruneQueue(ctx context.Context, pruneService UnstartedTxQueuePruner) ([]int64, error) {
	ret := _m.Called(ctx, pruneService)

	if len(ret) == 0 {
		panic("no return value specified for PruneQueue")
	}

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, UnstartedTxQueuePruner) ([]int64, error)); ok {
		return rf(ctx, pruneService)
	}
	if rf, ok := ret.Get(0).(func(context.Context, UnstartedTxQueuePruner) []int64); ok {
		r0 = rf(ctx, pruneService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, UnstartedTxQueuePruner) error); ok {
		r1 = rf(ctx, pruneService)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxStrategy_PruneQueue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PruneQueue'
type TxStrategy_PruneQueue_Call struct {
	*mock.Call
}

// PruneQueue is a helper method to define mock.On call
//   - ctx context.Context
//   - pruneService UnstartedTxQueuePruner
func (_e *TxStrategy_Expecter) PruneQueue(ctx interface{}, pruneService interface{}) *TxStrategy_PruneQueue_Call {
	return &TxStrategy_PruneQueue_Call{Call: _e.mock.On("PruneQueue", ctx, pruneService)}
}

func (_c *TxStrategy_PruneQueue_Call) Run(run func(ctx context.Context, pruneService UnstartedTxQueuePruner)) *TxStrategy_PruneQueue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(UnstartedTxQueuePruner))
	})
	return _c
}

func (_c *TxStrategy_PruneQueue_Call) Return(ids []int64, err error) *TxStrategy_PruneQueue_Call {
	_c.Call.Return(ids, err)
	return _c
}

func (_c *TxStrategy_PruneQueue_Call) RunAndReturn(run func(context.Context, UnstartedTxQueuePruner) ([]int64, error)) *TxStrategy_PruneQueue_Call {
	_c.Call.Return(run)
	return _c
}

// Subject provides a mock function with given fields:
func (_m *TxStrategy) Subject() uuid.NullUUID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Subject")
	}

	var r0 uuid.NullUUID
	if rf, ok := ret.Get(0).(func() uuid.NullUUID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uuid.NullUUID)
	}

	return r0
}

// TxStrategy_Subject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subject'
type TxStrategy_Subject_Call struct {
	*mock.Call
}

// Subject is a helper method to define mock.On call
func (_e *TxStrategy_Expecter) Subject() *TxStrategy_Subject_Call {
	return &TxStrategy_Subject_Call{Call: _e.mock.On("Subject")}
}

func (_c *TxStrategy_Subject_Call) Run(run func()) *TxStrategy_Subject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TxStrategy_Subject_Call) Return(_a0 uuid.NullUUID) *TxStrategy_Subject_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TxStrategy_Subject_Call) RunAndReturn(run func() uuid.NullUUID) *TxStrategy_Subject_Call {
	_c.Call.Return(run)
	return _c
}

// NewTxStrategy creates a new instance of TxStrategy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTxStrategy(t interface {
	mock.TestingT
	Cleanup(func())
}) *TxStrategy {
	mock := &TxStrategy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
