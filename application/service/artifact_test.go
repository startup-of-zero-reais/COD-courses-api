package service_test

import (
	"errors"
	"github.com/startup-of-zero-reais/COD-courses-api/application/service"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	mocks "github.com/startup-of-zero-reais/COD-courses-api/mocks/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/tests/entity_mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArtifactServiceAdd(t *testing.T) {
	preAddTest := func(returns ...interface{}) (domain.ArtifactRepository, domain.Artifact) {
		artifactSpy := *entity_mocks.ArtifactMock()

		repo := new(mocks.ArtifactRepository)
		repo.On("Create", artifactSpy).Return(returns...)

		return repo, artifactSpy
	}

	t.Run("should add artifact", func(t *testing.T) {
		artifactDuble := entity_mocks.ArtifactMock("mock-uuid")
		repo, artifactSpy := preAddTest(artifactDuble, nil)

		svc := service.NewArtifactService(repo)
		result, err := svc.Add(artifactSpy)

		require.Nil(t, err)
		require.NotEmpty(t, result)
		require.Equal(t, artifactSpy.LessonID, result.LessonID)
		require.Equal(t, artifactSpy.Link, result.Link)
		require.NotNil(t, result.ArtifactID)
	})
	t.Run("should not add artifact and return error", func(t *testing.T) {
		repo, artifactSpy := preAddTest(nil, errors.New("erro ao criar artefato"))

		svc := service.NewArtifactService(repo)
		result, err := svc.Add(artifactSpy)

		require.Nil(t, result)
		require.Error(t, err, "erro ao criar artefato")
	})
}
