package service_test

import (
	"github.com/startup-of-zero-reais/COD-courses-api/application/repository"
	"github.com/startup-of-zero-reais/COD-courses-api/application/service"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/tests/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

type Sut struct {
	svc domain.ArtifactService
}

func makeSut() *Sut {
	svc := &service.ArtifactServiceImpl{
		Repo: &repository.ArtifactRepositoryImpl{
			Db: &mocks.FakeArtifactDb{},
		},
	}

	sut := &Sut{
		svc: svc,
	}

	return sut
}

func TestArtifactServiceAdd(t *testing.T) {
	sut := makeSut()
	svc := sut.svc

	artifact := mocks.MockArtifact("", "")
	artifact.ArtifactID = ""
	require.Zero(t, artifact.ArtifactID)

	result, err := svc.Add(artifact)
	count, _ := svc.List(artifact.LessonID)

	require.Nil(t, err)
	require.NotEmpty(t, result)
	require.Len(t, count, 1)
}

func TestArtifactServiceList(t *testing.T) {
	sut := makeSut()
	svc := sut.svc

	artifact := mocks.MockArtifact("", "lesson-mock-uuid")
	_, err := svc.Add(artifact)
	require.Nil(t, err)

	result, err := svc.List("lesson-mock-uuid")
	require.Nil(t, err)
	require.NotEmpty(t, result)
	require.Len(t, result, 1)
}

func TestArtifactServiceRemove(t *testing.T) {
	sut := makeSut()
	svc := sut.svc

	artifactSpy := mocks.MockArtifact("", "")
	artifactSpy.ArtifactID = ""
	require.Zero(t, artifactSpy.ArtifactID)

	artifact, err := svc.Add(artifactSpy)
	require.Nil(t, err)
	require.NotNil(t, artifact.ArtifactID)

	err = svc.Remove(artifact.ArtifactID)
	require.Nil(t, err)
}
