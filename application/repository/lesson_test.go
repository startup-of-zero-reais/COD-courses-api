package repository_test

import (
	"github.com/google/uuid"
	"github.com/startup-of-zero-reais/COD-courses-api/application/repository"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	mocks "github.com/startup-of-zero-reais/COD-courses-api/mocks/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/tests/entity_mocks"
	"github.com/startup-of-zero-reais/COD-courses-api/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLessonRepositoryImpl_Create(t *testing.T) {
	preCreateTest := func(overrideMocks ...func(args mock.Arguments)) (domain.Db, domain.Lesson, string) {
		lessonSpy := *entity_mocks.LessonMock()

		var expected domain.Lesson
		expectedID := uuid.NewString()
		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Lesson)
			arg.LessonID = expectedID
			arg.SectionID = lessonSpy.SectionID
			arg.VideoSource = lessonSpy.VideoSource
			arg.DurationTotal = lessonSpy.DurationTotal
			arg.CreatedAt = lessonSpy.CreatedAt
			arg.UpdatedAt = lessonSpy.UpdatedAt

			if len(overrideMocks) > 0 {
				for _, override := range overrideMocks {
					override(args)
				}
			}
		}

		Db := new(mocks.Db)
		Db.On("Create", lessonSpy, &expected).Return().Run(mockResult)

		return Db, lessonSpy, expectedID
	}

	t.Run("should create an lesson", func(t *testing.T) {
		Db, lessonSpy, expectedID := preCreateTest()

		repo := repository.NewLessonRepository(Db)

		result, err := repo.Create(lessonSpy)

		require.Nil(t, err)
		require.NotNil(t, result)
		require.Zero(t, lessonSpy.LessonID)
		require.Equal(t, expectedID, result.LessonID)
		require.NotZero(t, result.LessonID)
		require.Equal(t, lessonSpy.SectionID, result.SectionID)
	})
	t.Run("should fail when no create", func(t *testing.T) {
		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Lesson)
			arg.LessonID = ""
		}

		Db, lessonSpy, _ := preCreateTest(mockResult)

		repo := repository.NewLessonRepository(Db)
		expected, err := repo.Create(lessonSpy)

		require.Nil(t, expected)
		require.NotNil(t, err)
		require.Error(t, err)
	})
}

func TestLessonRepositoryImpl_Get(t *testing.T) {
	t.Run("should get lessons", func(t *testing.T) {
		search := map[string]string{
			"section_id": "same-section-id",
		}
		pagination := map[string]string{
			"page":     "1",
			"per_page": "10",
		}

		var result []domain.Lesson
		var expected []domain.Lesson

		Db := new(mocks.Db)
		Db.On("Search", util.MergeMaps(search, pagination), &result).Return().Run(func(args mock.Arguments) {
			arg := args.Get(1).(*[]domain.Lesson)
			*arg = append(*arg,
				*entity_mocks.LessonMock(map[string]interface{}{"lesson_id": uuid.NewString(), "section_id": "same-section-id"}),
				*entity_mocks.LessonMock(map[string]interface{}{"lesson_id": uuid.NewString(), "section_id": "same-section-id"}),
			)
			expected = *arg
		})

		repo := repository.NewLessonRepository(Db)
		_, err := repo.Get(search, pagination)

		require.Nil(t, err)
		require.Len(t, expected, 2)
	})
}
