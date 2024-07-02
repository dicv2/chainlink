// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/types"
	mock "github.com/stretchr/testify/mock"
)

// Receiver is an autogenerated mock type for the Receiver type
type Receiver struct {
	mock.Mock
}

// Receive provides a mock function with given fields: ctx, msg
func (_m *Receiver) Receive(ctx context.Context, msg *types.MessageBody) {
	_m.Called(ctx, msg)
}

// NewReceiver creates a new instance of Receiver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReceiver(t interface {
	mock.TestingT
	Cleanup(func())
}) *Receiver {
	mock := &Receiver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
