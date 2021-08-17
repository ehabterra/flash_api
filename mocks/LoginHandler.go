// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	middleware "github.com/go-openapi/runtime/middleware"
	mock "github.com/stretchr/testify/mock"

	users "github.com/ehabterra/flash_api/api/restapi/operations/users"
)

// LoginHandler is an autogenerated mock type for the LoginHandler type
type LoginHandler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: _a0
func (_m *LoginHandler) Handle(_a0 users.LoginParams) middleware.Responder {
	ret := _m.Called(_a0)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(users.LoginParams) middleware.Responder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}