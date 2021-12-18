package mocks

import (
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	"github.com/stretchr/testify/mock"
)

type ArtifactRepositoryMock struct {
	mock.Mock
}

func (a *ArtifactRepositoryMock) Create(artifact domain.Artifact) (*domain.Artifact, error) {
	args := a.Called(artifact)
	return args.Get(0).(*domain.Artifact), args.Error(1)
}

func (a *ArtifactRepositoryMock) Save(artifact domain.Artifact) (*domain.Artifact, error) {
	args := a.Called(artifact)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Artifact), args.Error(1)
}

func (a *ArtifactRepositoryMock) Get(searchParam map[string]string) ([]domain.Artifact, error) {
	args := a.Called(searchParam)

	return args.Get(0).([]domain.Artifact), args.Error(1)
}

func (a *ArtifactRepositoryMock) Delete(artifactId string) error {
	a.Called(artifactId)
	return nil
}
