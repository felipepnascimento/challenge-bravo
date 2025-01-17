// Code generated by mockery v2.9.6. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// CurrencyController is an autogenerated mock type for the CurrencyController type
type CurrencyController struct {
	mock.Mock
}

// CreateCurrency provides a mock function with given fields: c
func (_m *CurrencyController) CreateCurrency(c *gin.Context) {
	_m.Called(c)
}

// DeleteCurrency provides a mock function with given fields: c
func (_m *CurrencyController) DeleteCurrency(c *gin.Context) {
	_m.Called(c)
}

// GetAllCurrencies provides a mock function with given fields: c
func (_m *CurrencyController) GetAllCurrencies(c *gin.Context) {
	_m.Called(c)
}

// GetCurrencyById provides a mock function with given fields: c
func (_m *CurrencyController) GetCurrencyById(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewCurrencyController interface {
	mock.TestingT
	Cleanup(func())
}

// NewCurrencyController creates a new instance of CurrencyController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCurrencyController(t mockConstructorTestingTNewCurrencyController) *CurrencyController {
	mock := &CurrencyController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
