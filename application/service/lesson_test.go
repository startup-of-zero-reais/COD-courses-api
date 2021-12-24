package service_test

import (
	"errors"
	"github.com/startup-of-zero-reais/COD-courses-api/application/service"
	mocks "github.com/startup-of-zero-reais/COD-courses-api/mocks/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLessonServiceImpl_Add(t *testing.T) {}

func TestLessonServiceImpl_Save(t *testing.T) {}

func TestLessonServiceImpl_ListBySection(t *testing.T) {}

func TestLessonServiceImpl_Get(t *testing.T) {}

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
