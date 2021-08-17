// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	constants "github.com/ehabterra/flash_api/internal/constants"
	mock "github.com/stretchr/testify/mock"

	models "github.com/ehabterra/flash_api/internal/models"

	time "time"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// AddAccountTransaction provides a mock function with given fields: transaction
func (_m *Database) AddAccountTransaction(transaction *models.AccountTransaction) error {
	ret := _m.Called(transaction)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.AccountTransaction) error); ok {
		r0 = rf(transaction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddRecipientTransaction provides a mock function with given fields: transaction
func (_m *Database) AddRecipientTransaction(transaction *models.RecipientTransaction) error {
	ret := _m.Called(transaction)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.RecipientTransaction) error); ok {
		r0 = rf(transaction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckAccountNumber provides a mock function with given fields: accountNumber
func (_m *Database) CheckAccountNumber(accountNumber string) (bool, error) {
	ret := _m.Called(accountNumber)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(accountNumber)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckTransactionLimits provides a mock function with given fields: id, transactionType, amount, limits
func (_m *Database) CheckTransactionLimits(id string, transactionType constants.TransactionType, amount int64, limits map[time.Duration]int64) (bool, error) {
	ret := _m.Called(id, transactionType, amount, limits)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, constants.TransactionType, int64, map[time.Duration]int64) bool); ok {
		r0 = rf(id, transactionType, amount, limits)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, constants.TransactionType, int64, map[time.Duration]int64) error); ok {
		r1 = rf(id, transactionType, amount, limits)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Connect provides a mock function with given fields: account
func (_m *Database) Connect(account *models.Account) error {
	ret := _m.Called(account)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Account) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBalance provides a mock function with given fields: id
func (_m *Database) GetBalance(id string) (float64, error) {
	ret := _m.Called(id)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsernameOrEmail provides a mock function with given fields: usernameOrEmail
func (_m *Database) GetUserByUsernameOrEmail(usernameOrEmail string) (*models.User, error) {
	ret := _m.Called(usernameOrEmail)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(usernameOrEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(usernameOrEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBalance provides a mock function with given fields: id, amount
func (_m *Database) UpdateBalance(id string, amount int64) error {
	ret := _m.Called(id, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(id, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}