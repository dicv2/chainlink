// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	pg "github.com/smartcontractkit/chainlink/core/services/pg"
	mock "github.com/stretchr/testify/mock"
)

// Subscription is an autogenerated mock type for the Subscription type
type Subscription struct {
	mock.Mock
}

// ChannelName provides a mock function with given fields:
func (_m *Subscription) ChannelName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Subscription) Close() {
	_m.Called()
}

// Events provides a mock function with given fields:
func (_m *Subscription) Events() <-chan pg.Event {
	ret := _m.Called()

	var r0 <-chan pg.Event
	if rf, ok := ret.Get(0).(func() <-chan pg.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan pg.Event)
		}
	}

	return r0
}

// InterestedIn provides a mock function with given fields: event
func (_m *Subscription) InterestedIn(event pg.Event) bool {
	ret := _m.Called(event)

	var r0 bool
	if rf, ok := ret.Get(0).(func(pg.Event) bool); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Send provides a mock function with given fields: event
func (_m *Subscription) Send(event pg.Event) {
	_m.Called(event)
}
