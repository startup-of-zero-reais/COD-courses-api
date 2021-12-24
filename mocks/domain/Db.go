// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/startup-of-zero-reais/COD-courses-api/domain"
	mock "github.com/stretchr/testify/mock"
)

// Db is an autogenerated mock type for the Db type
type Db struct {
	mock.Mock
}

// Connect provides a mock function with given fields:
func (_m *Db) Connect() {
	_m.Called()
}

// Create provides a mock function with given fields: entity, result
func (_m *Db) Create(entity interface{}, result domain.Result) {
	_m.Called(entity, result)
}

// Delete provides a mock function with given fields: param, result
func (_m *Db) Delete(param map[string]string, result domain.Result) bool {
	ret := _m.Called(param, result)

	var r0 bool
	if rf, ok := ret.Get(0).(func(map[string]string, domain.Result) bool); ok {
		r0 = rf(param, result)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Save provides a mock function with given fields: entity, result
func (_m *Db) Save(entity interface{}, result domain.Result) {
	_m.Called(entity, result)
}

// Search provides a mock function with given fields: param, result
func (_m *Db) Search(param map[string]string, result domain.Result) {
	_m.Called(param, result)
}

// TotalRows provides a mock function with given fields:
func (_m *Db) TotalRows() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}
