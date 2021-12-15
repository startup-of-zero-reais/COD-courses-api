package repository_test

import (
	"github.com/startup-of-zero-reais/COD-courses-api/application/repository"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/tests/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

type Sut struct {
	Db   *mocks.DatabaseUtil
	repo domain.ArtifactRepository
}

func makeSut() *Sut {
	Db := mocks.SetupTest()
	Db.Env = "testing"

	return &Sut{
		Db: Db,
		repo: &repository.ArtifactRepositoryImpl{
			Db: Db.Database,
		},
	}
}

// TODO : FIX NIL POINTER
func TestArtifactRepositoryCreate(t *testing.T) {
	sut := makeSut()
	sut.Db.BeforeTests()
	defer sut.Db.AfterTests()

	artifactSpy := mocks.MockArtifact("", "")
	artifactSpy.ArtifactID = ""
	require.Zero(t, artifactSpy.ArtifactID)
	require.NotZero(t, artifactSpy.LessonID)

	artifact, err := sut.repo.Create(artifactSpy)
	require.Nil(t, err)
	require.NotZero(t, artifact.ArtifactID)
	require.NotZero(t, artifact.LessonID)
	require.NotZero(t, artifact.Link)

	list, err := sut.repo.Get(map[string]string{"lesson_id": artifact.LessonID})
	require.Nil(t, err)
	require.Len(t, list, 1)
}

func TestArtifactRepositorySave(t *testing.T) {
	sut := makeSut()
	sut.Db.BeforeTests()
	defer sut.Db.AfterTests()

	artifactSpy := mocks.MockArtifact("", "lesson-1")
	artifact, err := sut.repo.Save(artifactSpy)
	artifactSpy.ArtifactID = artifact.ArtifactID

	require.Nil(t, err)
	require.NotEmpty(t, artifact)
	require.Equal(t, artifact.LessonID, "lesson-1")

	find, err := sut.repo.Get(map[string]string{"artifact_id": artifact.ArtifactID})
	require.Nil(t, err)
	require.Len(t, find, 1)

	artifactSpy.LessonID = "lesson-2"
	updated, err := sut.repo.Save(artifactSpy)
	require.Nil(t, err)
	require.NotEmpty(t, updated)
	require.Equal(t, updated.ArtifactID, artifact.ArtifactID)
	require.Equal(t, updated.LessonID, "lesson-2")
}

func TestArtifactRepositoryGet(t *testing.T) {
	sut := makeSut()
	sut.Db.BeforeTests()
	defer sut.Db.AfterTests()

	artifactSpy := mocks.MockArtifact("", "lesson-1")
	_, err := sut.repo.Create(artifactSpy)
	require.Nil(t, err)

	get, err := sut.repo.Get(map[string]string{"lesson_id": "lesson-1"})
	require.Nil(t, err)
	require.NotEmpty(t, get)
	require.Equal(t, get[0].LessonID, "lesson-1")
}

func TestArtifactRepositoryDelete(t *testing.T) {
	sut := makeSut()
	sut.Db.BeforeTests()
	defer sut.Db.AfterTests()

	err := sut.repo.Delete("artifact-mock-uuid")
	require.Error(t, err, "ocorreu um erro ao deletar artefato da base de dados")

	artifactSpy := mocks.MockArtifact("artifact-mock-uuid", "lesson-1")
	artifact, err := sut.repo.Create(artifactSpy)
	require.Nil(t, err)

	err = sut.repo.Delete(artifact.ArtifactID)
	require.Nil(t, err)

	find, err := sut.repo.Get(map[string]string{"artifact_id": artifact.ArtifactID})
	require.Nil(t, err)
	require.Len(t, find, 0)
	require.Empty(t, find)
}
