// Code generated by mockery (devel). DO NOT EDIT.

package q3

import (
	mock "github.com/stretchr/testify/mock"
)

// MessageQueue is an autogenerated mock type for the MessageQueue type
type MockMessageQueue struct {
	mock.Mock
}

// Publish provides a mock function with given fields: _a0, _a1
func (_m *MockMessageQueue) Publish(_a0 MessageName, _a1 map[string]interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(MessageName, map[string]interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
