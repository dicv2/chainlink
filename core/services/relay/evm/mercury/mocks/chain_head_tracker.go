// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	client "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/v2/core/chains/evm/headmanager/types"
)

// ChainHeadTracker is an autogenerated mock type for the ChainHeadTracker type
type ChainHeadTracker struct {
	mock.Mock
}

// Client provides a mock function with given fields:
func (_m *ChainHeadTracker) Client() client.Client {
	ret := _m.Called()

	var r0 client.Client
	if rf, ok := ret.Get(0).(func() client.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Client)
		}
	}

	return r0
}

// HeadTracker provides a mock function with given fields:
func (_m *ChainHeadTracker) HeadTracker() types.Tracker {
	ret := _m.Called()

	var r0 types.Tracker
	if rf, ok := ret.Get(0).(func() types.Tracker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Tracker)
		}
	}

	return r0
}

type mockConstructorTestingTNewChainHeadTracker interface {
	mock.TestingT
	Cleanup(func())
}

// NewChainHeadTracker creates a new instance of ChainHeadTracker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChainHeadTracker(t mockConstructorTestingTNewChainHeadTracker) *ChainHeadTracker {
	mock := &ChainHeadTracker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
