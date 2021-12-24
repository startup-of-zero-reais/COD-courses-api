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

func TestLessonServiceImpl_Add(t *testing.T) {}

func TestLessonServiceImpl_Save(t *testing.T) {}

func TestLessonServiceImpl_ListBySection(t *testing.T) {
	preListTests := func(returns []domain.Lesson, errReturn error) (domain.LessonService, map[string]string, map[string]string) {
		repo := new(mocks.LessonRepository)
		search := map[string]string{
			"section_id": "section-id",
		}
		query := map[string]string{
			"page":     "1",
			"per_page": "10",
		}

		repo.On("Get", search, query).Return(returns, errReturn)

		svc := service.NewLessonService(repo)
		return svc, search, query
	}

	t.Run("should list lessons of a section", func(t *testing.T) {
		returns := []domain.Lesson{
			*entity_mocks.LessonMock(map[string]interface{}{"section_id": "section-id"}),
			*entity_mocks.LessonMock(map[string]interface{}{"section_id": "section-id"}),
		}
		svc, _, query := preListTests(returns, nil)

		expected, err := svc.ListBySection("section-id", query)

		require.Nil(t, err)
		require.Len(t, expected, 2)
	})
	t.Run("should return an empty slice", func(t *testing.T) {
		var returns []domain.Lesson
		svc, _, query := preListTests(returns, nil)

		expected, err := svc.ListBySection("section-id", query)

		require.Nil(t, err)
		require.Empty(t, expected)
	})
	t.Run("should fail if section not exists", func(t *testing.T) {
		svc, _, query := preListTests(nil, errors.New("não foi possível recuperar aulas de uma seção inexistente"))

		expected, err := svc.ListBySection("section-id", query)

		require.Nil(t, expected)
		require.EqualError(t, err, "não foi possível recuperar aulas de uma seção inexistente")
	})
}

func TestLessonServiceImpl_Get(t *testing.T) {
	preGetTests := func(argReturns []domain.Lesson, errReturn error) domain.LessonRepository {
		search := map[string]string{
			"lesson_id": "lesson-id",
		}
		pagination := map[string]string{
			"page":     "1",
			"per_page": "1",
		}

		repo := new(mocks.LessonRepository)

		repo.On("Get", search, pagination).Return(argReturns, errReturn)

		return repo
	}

	t.Run("should get a single lesson", func(t *testing.T) {
		lessonSpy := entity_mocks.LessonMock()

		repo := preGetTests([]domain.Lesson{*lessonSpy}, nil)

		svc := service.NewLessonService(repo)
		expected, err := svc.Get("lesson-id")

		require.Nil(t, err)
		require.NotNil(t, expected)
	})
	t.Run("should fail when lesson does not exists", func(t *testing.T) {
		repo := preGetTests(nil, errors.New("erro ao buscar a aula"))

		svc := service.NewLessonService(repo)
		expected, err := svc.Get("lesson-id")

		require.Nil(t, expected)
		require.EqualError(t, err, "erro ao buscar a aula")
	})
	t.Run("should fail if has no results", func(t *testing.T) {
		repo := preGetTests([]domain.Lesson{}, nil)

		svc := service.NewLessonService(repo)
		expected, err := svc.Get("lesson-id")

		require.Nil(t, expected)
		require.EqualError(t, err, "nenhuma aula encontrada")
	})
	t.Run("should fail if has more than a single result", func(t *testing.T) {
		repo := preGetTests([]domain.Lesson{
			*entity_mocks.LessonMock(map[string]interface{}{"lesson_id": "lesson-id-1"}),
			*entity_mocks.LessonMock(map[string]interface{}{"lesson_id": "lesson-id-2"}),
		}, nil)

		svc := service.NewLessonService(repo)
		expected, err := svc.Get("lesson-id")

		require.Nil(t, expected)
		require.EqualError(t, err, "mais aulas encontradas. contate o administrador do sistema")
	})
}

func TestLessonServiceImpl_Delete(t *testing.T) {
	t.Run("should delete a lesson", func(t *testing.T) {
		repo := new(mocks.LessonRepository)

		repo.On("Delete", "lesson-id").Return(nil)

		svc := service.NewLessonService(repo)
		err := svc.Delete("lesson-id")

		require.Nil(t, err)
	})
	t.Run("should fail on delete", func(t *testing.T) {
		repo := new(mocks.LessonRepository)

		repo.On("Delete", "lesson-id").Return(errors.New("erro ao deletar aula"))

		svc := service.NewLessonService(repo)
		err := svc.Delete("lesson-id")

		require.EqualError(t, err, "erro ao deletar aula")
	})
}
