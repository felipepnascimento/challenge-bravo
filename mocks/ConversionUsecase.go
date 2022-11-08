// Code generated by mockery v2.9.6. DO NOT EDIT.

package mocks

import (
	entities "github.com/felipepnascimento/challenge-bravo-flp/entities"
	mock "github.com/stretchr/testify/mock"
)

// ConversionUsecase is an autogenerated mock type for the ConversionUsecase type
type ConversionUsecase struct {
	mock.Mock
}

// CreateConversion provides a mock function with given fields: conversion
func (_m *ConversionUsecase) CreateConversion(conversion *entities.Conversion) error {
	ret := _m.Called(conversion)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Conversion) error); ok {
		r0 = rf(conversion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewConversionUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewConversionUsecase creates a new instance of ConversionUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConversionUsecase(t mockConstructorTestingTNewConversionUsecase) *ConversionUsecase {
	mock := &ConversionUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
