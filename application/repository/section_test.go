package repository_test

import (
	"github.com/google/uuid"
	"github.com/startup-of-zero-reais/COD-courses-api/application/repository"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	mocks "github.com/startup-of-zero-reais/COD-courses-api/mocks/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/tests/entity_mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSectionRepositoryImpl_Create(t *testing.T) {
	preCreateTest := func(result *domain.Section, overrides ...func(args mock.Arguments)) domain.SectionRepository {

		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Section)
			if result.SectionID != "" {
				arg.SectionID = result.SectionID
			} else {
				arg.SectionID = uuid.NewString()
			}
			arg.ModuleID = result.ModuleID
			arg.Label = result.Label
			arg.Lessons = result.Lessons

			if len(overrides) > 0 {
				for _, override := range overrides {
					override(args)
				}
			}
		}

		var r domain.Section
		Db := new(mocks.Db)
		Db.On("Create", *result, &r).Return().Run(mockResult)

		svc := repository.NewSectionRepository(Db)

		return svc
	}

	t.Run("should create a section", func(t *testing.T) {
		sectionSpy := entity_mocks.SectionMock(map[string]interface{}{"section_id": "-"})
		svc := preCreateTest(sectionSpy)

		expected, err := svc.Create(*sectionSpy)

		require.Nil(t, err)
		require.Zero(t, sectionSpy.SectionID)
		require.NotNil(t, expected)
	})
	t.Run("should fail on create a section", func(t *testing.T) {
		sectionSpy := entity_mocks.SectionMock(map[string]interface{}{
			"section_id": "-",
		})

		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Section)
			arg.SectionID = ""
			arg = nil
		}

		svc := preCreateTest(sectionSpy, mockResult)

		expected, err := svc.Create(*sectionSpy)

		require.Nil(t, expected)
		require.NotNil(t, err)
		require.EqualError(t, err, "erro ao criar seção")
	})
}

func TestSectionRepositoryImpl_Save(t *testing.T) {
	t.Run("should save a section", func(t *testing.T) {

	})
	t.Run("should fail on save", func(t *testing.T) {

	})
}

func TestSectionRepositoryImpl_Get(t *testing.T) {

}

func TestSectionRepositoryImpl_Delete(t *testing.T) {

}

func TestSectionRepositoryImpl_Count(t *testing.T) {

}
