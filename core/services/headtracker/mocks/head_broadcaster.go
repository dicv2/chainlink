// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	context "context"

	logger "github.com/smartcontractkit/chainlink/core/logger"
	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	types "github.com/smartcontractkit/chainlink/core/services/headtracker/types"
)

// HeadBroadcaster is an autogenerated mock type for the HeadBroadcaster type
type HeadBroadcaster struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *HeadBroadcaster) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Connect provides a mock function with given fields: head
func (_m *HeadBroadcaster) Connect(head *models.Head) error {
	ret := _m.Called(head)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Head) error); ok {
		r0 = rf(head)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Healthy provides a mock function with given fields:
func (_m *HeadBroadcaster) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *HeadBroadcaster) OnNewLongestChain(ctx context.Context, head models.Head) {
	_m.Called(ctx, head)
}

// Ready provides a mock function with given fields:
func (_m *HeadBroadcaster) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLogger provides a mock function with given fields: _a0
func (_m *HeadBroadcaster) SetLogger(_a0 logger.Logger) {
	_m.Called(_a0)
}

// Start provides a mock function with given fields:
func (_m *HeadBroadcaster) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: callback
func (_m *HeadBroadcaster) Subscribe(callback types.HeadTrackable) func() {
	ret := _m.Called(callback)

	var r0 func()
	if rf, ok := ret.Get(0).(func(types.HeadTrackable) func()); ok {
		r0 = rf(callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func())
		}
	}

	return r0
}
