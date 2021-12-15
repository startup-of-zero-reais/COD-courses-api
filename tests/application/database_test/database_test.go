package database_test

import (
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/tests/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDatabaseCreate(t *testing.T) {
	db := mocks.SetupTest()
	db.BeforeTests()
	defer db.AfterTests()

	artifact := &domain.Artifact{
		ArtifactID: "",
		LessonID:   "lesson-uuid",
		Link:       "http:link",
	}
	require.Zero(t, artifact.ArtifactID)

	var result domain.Artifact
	db.Create(artifact, &result)

	require.NotZero(t, result.ArtifactID)
	require.NotEmpty(t, result)
}

func TestDatabaseSave(t *testing.T) {
	db := mocks.SetupTest()
	db.BeforeTests()
	defer db.AfterTests()

	artifact := &domain.Artifact{
		LessonID: "lesson-save",
		Link:     "save-link",
	}

	var result domain.Artifact
	db.Save(artifact, &result)
	require.NotEmpty(t, result)
	require.NotNil(t, result.ArtifactID)

	artifact.LessonID = "lesson-update"
	db.Save(artifact, &result)
	require.Equal(t, artifact.LessonID, result.LessonID)
	require.Equal(t, result.LessonID, "lesson-update")
}

func TestDatabaseSearch(t *testing.T) {
	db := mocks.SetupTest()
	db.BeforeTests()
	defer db.AfterTests()

	querySearch := map[string]string{
		"lesson_id": "lesson-mock",
	}
	var result domain.Artifact
	db.Search(querySearch, &result)
	require.Empty(t, result)

	db.Create(&domain.Artifact{
		LessonID: "lesson-mock",
		Link:     "link",
	}, &(struct{}{}))

	db.Search(querySearch, &result)
	require.NotEmpty(t, result)
	require.Equal(t, result.LessonID, "lesson-mock")
}

func TestDatabaseDelete(t *testing.T) {
	db := mocks.SetupTest()
	db.BeforeTests()
	defer db.AfterTests()

	deleteQuery := map[string]string{
		"lesson_id": "lesson-mock",
	}
	wasDeleted := db.Delete(deleteQuery, domain.Artifact{})
	require.False(t, wasDeleted)

	db.Create(&domain.Artifact{
		LessonID: "lesson-mock",
		Link:     "link",
	}, &(struct{}{}))

	wasDeleted = db.Delete(deleteQuery, domain.Artifact{})
	require.True(t, wasDeleted)
	var result domain.Artifact
	db.Search(deleteQuery, &result)
	require.Empty(t, result)
}
