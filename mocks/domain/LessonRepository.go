// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/startup-of-zero-reais/COD-courses-api/domain"
	mock "github.com/stretchr/testify/mock"
)

// LessonRepository is an autogenerated mock type for the LessonRepository type
type LessonRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields:
func (_m *LessonRepository) Count() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// Create provides a mock function with given fields: lesson
func (_m *LessonRepository) Create(lesson domain.Lesson) (*domain.Lesson, error) {
	ret := _m.Called(lesson)

	var r0 *domain.Lesson
	if rf, ok := ret.Get(0).(func(domain.Lesson) *domain.Lesson); ok {
		r0 = rf(lesson)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Lesson)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Lesson) error); ok {
		r1 = rf(lesson)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: lessonID
func (_m *LessonRepository) Delete(lessonID string) error {
	ret := _m.Called(lessonID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(lessonID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: search, pagination
func (_m *LessonRepository) Get(search map[string]string, pagination map[string]string) ([]domain.Lesson, error) {
	ret := _m.Called(search, pagination)

	var r0 []domain.Lesson
	if rf, ok := ret.Get(0).(func(map[string]string, map[string]string) []domain.Lesson); ok {
		r0 = rf(search, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Lesson)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]string, map[string]string) error); ok {
		r1 = rf(search, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: lesson
func (_m *LessonRepository) Save(lesson domain.Lesson) (*domain.Lesson, error) {
	ret := _m.Called(lesson)

	var r0 *domain.Lesson
	if rf, ok := ret.Get(0).(func(domain.Lesson) *domain.Lesson); ok {
		r0 = rf(lesson)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Lesson)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Lesson) error); ok {
		r1 = rf(lesson)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
