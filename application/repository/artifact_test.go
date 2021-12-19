package repository_test

import (
	"github.com/startup-of-zero-reais/COD-courses-api/application/repository"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	mocks "github.com/startup-of-zero-reais/COD-courses-api/mocks/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/tests/entity_mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArtifactRepositoryCreate(t *testing.T) {
	preCreateTest := func(overrideResults ...func(args mock.Arguments)) (domain.Db, domain.Artifact, string) {
		Db := new(mocks.Db)

		artifactSpy := *entity_mocks.ArtifactMock("", "", "")

		var expected domain.Artifact
		expectedID := "mock-uuid"

		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Artifact)

			arg.ArtifactID = expectedID
			arg.LessonID = artifactSpy.LessonID
			arg.Link = artifactSpy.Link

			if len(overrideResults) > 0 {
				for _, override := range overrideResults {
					override(args)
				}
			}
		}

		Db.On("Create", artifactSpy, &expected).Return().Run(mockResult)

		return Db, artifactSpy, expectedID
	}

	t.Run("should create an artifact", func(t *testing.T) {
		Db, artifactSpy, expectedID := preCreateTest()

		repo := repository.NewArtifactRepository(Db)
		result, err := repo.Create(artifactSpy)

		require.Nil(t, err)
		require.Zero(t, artifactSpy.ArtifactID)
		require.NotZero(t, result.ArtifactID)
		require.Equal(t, result.ArtifactID, expectedID)
	})
	t.Run("should return error when creation fails", func(t *testing.T) {
		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Artifact)
			arg.ArtifactID = ""
			arg = nil
		}
		Db, artifactSpy, _ := preCreateTest(mockResult)

		repo := repository.NewArtifactRepository(Db)
		result, err := repo.Create(artifactSpy)

		require.Nil(t, result)
		require.Error(t, err, "falha ao criar artefato na base de dados")
		require.NotEmpty(t, artifactSpy)
	})
}

func TestArtifactRepositorySave(t *testing.T) {
	preSaveTest := func(overrideResult ...func(args mock.Arguments)) (domain.Db, *domain.Artifact) {
		Db := new(mocks.Db)

		artifactSpy := *entity_mocks.ArtifactMock("mock-uuid", "", "")

		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Artifact)
			arg.ArtifactID = artifactSpy.ArtifactID
			arg.LessonID = artifactSpy.LessonID
			arg.Link = artifactSpy.Link

			if len(overrideResult) > 0 {
				for _, override := range overrideResult {
					override(args)
				}
			}
		}

		Db.On("Save", artifactSpy, &domain.Artifact{}).Return().Run(mockResult)

		return Db, &artifactSpy
	}

	t.Run("should save an artifact", func(t *testing.T) {
		Db, artifactSpy := preSaveTest()

		repo := repository.NewArtifactRepository(Db)
		result, err := repo.Save(*artifactSpy)

		require.Nil(t, err)
		require.NotEmpty(t, result)
		require.Equal(t, artifactSpy.ArtifactID, result.ArtifactID)
	})
	t.Run("should return an error if artifacts does not exists", func(t *testing.T) {
		Db, artifactSpy := preSaveTest()
		artifactSpy.ArtifactID = ""

		repo := repository.NewArtifactRepository(Db)
		result, err := repo.Save(*artifactSpy)

		require.Nil(t, result)
		require.Error(t, err, "artefato sem referência no banco")
	})
	t.Run("should return error if db cannot save artifact", func(t *testing.T) {
		overrideResult := func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Artifact)
			arg.ArtifactID = ""
			arg = nil
		}

		Db, artifactSpy := preSaveTest(overrideResult)

		repo := repository.NewArtifactRepository(Db)
		result, err := repo.Save(*artifactSpy)

		require.Nil(t, result)
		require.Error(t, err, "ocorreu algum erro ao salvar artefato")
	})
}

func TestArtifactRepositoryGet(t *testing.T) {
	preGetTests := func(searchParam map[string]string, overrideResults ...func(args mock.Arguments)) (domain.Db, *[]domain.Artifact) {
		Db := new(mocks.Db)

		var expected []domain.Artifact
		mockResult := func(args mock.Arguments) {
			artifactSpy1 := *entity_mocks.ArtifactMock("artifact-1", "lesson-mock", "")
			artifactSpy2 := *entity_mocks.ArtifactMock("artifact-2", "lesson-mock", "")

			arg := args.Get(1).(*[]domain.Artifact)
			*arg = append(*arg, artifactSpy1, artifactSpy2)
			expected = *arg

			if len(overrideResults) > 0 {
				for _, override := range overrideResults {
					override(args)
				}
			}
		}

		var artifacts []domain.Artifact
		Db.On("Search", searchParam, &artifacts).Return().Run(mockResult)

		return Db, &expected
	}

	t.Run("should return an artifact", func(t *testing.T) {
		searchParam := map[string]string{
			"lesson_id": "lesson-mock",
		}

		Db, expected := preGetTests(searchParam)

		repo := repository.NewArtifactRepository(Db)
		result, err := repo.Get(searchParam)

		require.Nil(t, err)
		require.NotEmpty(t, result)
		require.Len(t, result, 2)
		require.Equal(t, len(*expected), len(result))
	})
	t.Run("should return an empty slice", func(t *testing.T) {
		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*[]domain.Artifact)
			*arg = []domain.Artifact{}
		}
		searchParam := map[string]string{
			"artifact_id": "void-uuid",
		}
		Db, _ := preGetTests(searchParam, mockResult)

		repo := repository.NewArtifactRepository(Db)
		result, err := repo.Get(searchParam)

		require.Nil(t, err)
		require.Empty(t, result)
	})
}

func TestArtifactRepositoryDelete(t *testing.T) {

	t.Run("should delete an artifact", func(t *testing.T) {
		Db := new(mocks.Db)

		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Artifact)
			arg.ArtifactID = ""
			arg = nil
		}
		Db.On("Delete", map[string]string{"artifact_id": "mock-uuid"}, &domain.Artifact{}).Return(true).Run(mockResult)

		repo := repository.NewArtifactRepository(Db)
		result := repo.Delete("mock-uuid")

		require.Nil(t, result)
	})
	t.Run("should not delete an artifact and return error", func(t *testing.T) {
		Db := new(mocks.Db)

		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Artifact)
			arg.ArtifactID = ""
			arg = nil
		}
		Db.On("Delete", map[string]string{"artifact_id": "mock-uuid"}, &domain.Artifact{}).Return(false).Run(mockResult)

		repo := repository.NewArtifactRepository(Db)
		result := repo.Delete("mock-uuid")

		require.Error(t, result, "ocorreu um erro ao deletar artefato da base de dados")
	})
}
