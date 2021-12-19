package service_test

import (
	"github.com/startup-of-zero-reais/COD-courses-api/application/service"
	mocks "github.com/startup-of-zero-reais/COD-courses-api/mocks/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/tests/entity_mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArtifactServiceAdd(t *testing.T) {
	t.Run("should add artifact", func(t *testing.T) {
		artifactSpy := *entity_mocks.ArtifactMock()
		artifactDuble := entity_mocks.ArtifactMock("mock-uuid")

		repo := new(mocks.ArtifactRepository)
		repo.On("Create", artifactSpy).Return(artifactDuble, nil)

		svc := service.NewArtifactService(repo)
		result, err := svc.Add(artifactSpy)

		require.Nil(t, err)
		require.NotEmpty(t, result)
		require.Equal(t, artifactSpy.LessonID, result.LessonID)
		require.Equal(t, artifactSpy.Link, result.Link)
		require.NotNil(t, result.ArtifactID)
	})
}
