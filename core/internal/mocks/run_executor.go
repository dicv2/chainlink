// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/satori/go.uuid"
)

// RunExecutor is an autogenerated mock type for the RunExecutor type
type RunExecutor struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *RunExecutor) Execute(_a0 uuid.UUID) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
