// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/startup-of-zero-reais/COD-courses-api/domain"
	mock "github.com/stretchr/testify/mock"
)

// ArtifactRepository is an autogenerated mock type for the ArtifactRepository type
type ArtifactRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields:
func (_m *ArtifactRepository) Count() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// Create provides a mock function with given fields: artifact
func (_m *ArtifactRepository) Create(artifact domain.Artifact) (*domain.Artifact, error) {
	ret := _m.Called(artifact)

	var r0 *domain.Artifact
	if rf, ok := ret.Get(0).(func(domain.Artifact) *domain.Artifact); ok {
		r0 = rf(artifact)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Artifact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Artifact) error); ok {
		r1 = rf(artifact)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: artifactId
func (_m *ArtifactRepository) Delete(artifactId string) error {
	ret := _m.Called(artifactId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(artifactId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: searchParam, pagination
func (_m *ArtifactRepository) Get(searchParam map[string]string, pagination map[string]string) ([]domain.Artifact, error) {
	ret := _m.Called(searchParam, pagination)

	var r0 []domain.Artifact
	if rf, ok := ret.Get(0).(func(map[string]string, map[string]string) []domain.Artifact); ok {
		r0 = rf(searchParam, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Artifact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]string, map[string]string) error); ok {
		r1 = rf(searchParam, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: artifact
func (_m *ArtifactRepository) Save(artifact domain.Artifact) (*domain.Artifact, error) {
	ret := _m.Called(artifact)

	var r0 *domain.Artifact
	if rf, ok := ret.Get(0).(func(domain.Artifact) *domain.Artifact); ok {
		r0 = rf(artifact)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Artifact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Artifact) error); ok {
		r1 = rf(artifact)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
