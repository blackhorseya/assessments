// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/amazingtalker/internal/pkg/contextx"

	mock "github.com/stretchr/testify/mock"
)

// IBiz is an autogenerated mock type for the IBiz type
type IBiz struct {
	mock.Mock
}

// Liveness provides a mock function with given fields: ctx
func (_m *IBiz) Liveness(ctx contextx.Contextx) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(contextx.Contextx) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Readiness provides a mock function with given fields: ctx
func (_m *IBiz) Readiness(ctx contextx.Contextx) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(contextx.Contextx) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
