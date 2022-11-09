// Code generated by mockery v2.9.6. DO NOT EDIT.

package mocks

import (
	models "github.com/felipepnascimento/challenge-bravo-flp/models"
	mock "github.com/stretchr/testify/mock"
)

// CurrencyRepository is an autogenerated mock type for the CurrencyRepository type
type CurrencyRepository struct {
	mock.Mock
}

// CreateCurrency provides a mock function with given fields: currency
func (_m *CurrencyRepository) CreateCurrency(currency *models.Currency) error {
	ret := _m.Called(currency)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Currency) error); ok {
		r0 = rf(currency)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCurrency provides a mock function with given fields: id
func (_m *CurrencyRepository) DeleteCurrency(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCurrencies provides a mock function with given fields:
func (_m *CurrencyRepository) GetAllCurrencies() (*[]models.Currency, error) {
	ret := _m.Called()

	var r0 *[]models.Currency
	if rf, ok := ret.Get(0).(func() *[]models.Currency); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Currency)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrencyById provides a mock function with given fields: id
func (_m *CurrencyRepository) GetCurrencyById(id int) (*models.Currency, error) {
	ret := _m.Called(id)

	var r0 *models.Currency
	if rf, ok := ret.Get(0).(func(int) *models.Currency); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Currency)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrencyByKey provides a mock function with given fields: key
func (_m *CurrencyRepository) GetCurrencyByKey(key string) (*models.Currency, error) {
	ret := _m.Called(key)

	var r0 *models.Currency
	if rf, ok := ret.Get(0).(func(string) *models.Currency); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Currency)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCurrencyRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCurrencyRepository creates a new instance of CurrencyRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCurrencyRepository(t mockConstructorTestingTNewCurrencyRepository) *CurrencyRepository {
	mock := &CurrencyRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
