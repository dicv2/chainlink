// Code generated by mockery v2.42.2. DO NOT EDIT.

package client

import (
	context "context"

	types "github.com/smartcontractkit/chainlink/v2/common/types"
	mock "github.com/stretchr/testify/mock"
)

// mockNodeClient is an autogenerated mock type for the NodeClient type
type mockNodeClient[CHAIN_ID types.ID, HEAD Head] struct {
	mock.Mock
}

// ChainID provides a mock function with given fields: ctx
func (_m *mockNodeClient[CHAIN_ID, HEAD]) ChainID(ctx context.Context) (CHAIN_ID, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ChainID")
	}

	var r0 CHAIN_ID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (CHAIN_ID, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) CHAIN_ID); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(CHAIN_ID)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientVersion provides a mock function with given fields: _a0
func (_m *mockNodeClient[CHAIN_ID, HEAD]) ClientVersion(_a0 context.Context) (string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ClientVersion")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *mockNodeClient[CHAIN_ID, HEAD]) Close() {
	_m.Called()
}

// Dial provides a mock function with given fields: ctx
func (_m *mockNodeClient[CHAIN_ID, HEAD]) Dial(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Dial")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DialHTTP provides a mock function with given fields:
func (_m *mockNodeClient[CHAIN_ID, HEAD]) DialHTTP() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DialHTTP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsSyncing provides a mock function with given fields: ctx
func (_m *mockNodeClient[CHAIN_ID, HEAD]) IsSyncing(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for IsSyncing")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (bool, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestFinalizedBlock provides a mock function with given fields: ctx
func (_m *mockNodeClient[CHAIN_ID, HEAD]) LatestFinalizedBlock(ctx context.Context) (HEAD, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LatestFinalizedBlock")
	}

	var r0 HEAD
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (HEAD, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) HEAD); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(HEAD)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAliveLoopSub provides a mock function with given fields: _a0
func (_m *mockNodeClient[CHAIN_ID, HEAD]) SetAliveLoopSub(_a0 types.Subscription) {
	_m.Called(_a0)
}

// SubscribeNewHead provides a mock function with given fields: ctx, channel
func (_m *mockNodeClient[CHAIN_ID, HEAD]) SubscribeNewHead(ctx context.Context, channel chan<- HEAD) (types.Subscription, error) {
	ret := _m.Called(ctx, channel)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeNewHead")
	}

	var r0 types.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, chan<- HEAD) (types.Subscription, error)); ok {
		return rf(ctx, channel)
	}
	if rf, ok := ret.Get(0).(func(context.Context, chan<- HEAD) types.Subscription); ok {
		r0 = rf(ctx, channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, chan<- HEAD) error); ok {
		r1 = rf(ctx, channel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribersCount provides a mock function with given fields:
func (_m *mockNodeClient[CHAIN_ID, HEAD]) SubscribersCount() int32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SubscribersCount")
	}

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// UnsubscribeAllExceptAliveLoop provides a mock function with given fields:
func (_m *mockNodeClient[CHAIN_ID, HEAD]) UnsubscribeAllExceptAliveLoop() {
	_m.Called()
}

// newMockNodeClient creates a new instance of mockNodeClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockNodeClient[CHAIN_ID types.ID, HEAD Head](t interface {
	mock.TestingT
	Cleanup(func())
}) *mockNodeClient[CHAIN_ID, HEAD] {
	mock := &mockNodeClient[CHAIN_ID, HEAD]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
