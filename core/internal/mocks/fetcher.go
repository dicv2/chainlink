// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	decimal "github.com/shopspring/decimal"

	mock "github.com/stretchr/testify/mock"
)

// Fetcher is an autogenerated mock type for the Fetcher type
type Fetcher struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: _a0, _a1
func (_m *Fetcher) Fetch(_a0 context.Context, _a1 map[string]interface{}) (decimal.Decimal, error) {
	ret := _m.Called(_a0, _a1)

	var r0 decimal.Decimal
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) decimal.Decimal); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(decimal.Decimal)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
