// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	middleware "github.com/go-openapi/runtime/middleware"
	mock "github.com/stretchr/testify/mock"

	models "github.com/ehabterra/flash_api/api/models"

	users "github.com/ehabterra/flash_api/api/restapi/operations/users"
)

// SendHandler is an autogenerated mock type for the SendHandler type
type SendHandler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: _a0, _a1
func (_m *SendHandler) Handle(_a0 users.SendParams, _a1 *models.Principle) middleware.Responder {
	ret := _m.Called(_a0, _a1)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(users.SendParams, *models.Principle) middleware.Responder); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}
