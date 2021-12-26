package repository_test

import (
	"github.com/google/uuid"
	"github.com/startup-of-zero-reais/COD-courses-api/application/repository"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	mocks "github.com/startup-of-zero-reais/COD-courses-api/mocks/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/tests/entity_mocks"
	"github.com/stretchr/testify/assert"
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

		repo := repository.NewSectionRepository(Db)

		return repo
	}

	t.Run("should create a section", func(t *testing.T) {
		sectionSpy := entity_mocks.SectionMock(map[string]interface{}{"section_id": "-"})
		repo := preCreateTest(sectionSpy)

		expected, err := repo.Create(*sectionSpy)

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

		repo := preCreateTest(sectionSpy, mockResult)

		expected, err := repo.Create(*sectionSpy)

		require.Nil(t, expected)
		require.NotNil(t, err)
		require.EqualError(t, err, "erro ao criar seção")
	})
}

func TestSectionRepositoryImpl_Save(t *testing.T) {
	preSaveTest := func(spy *domain.Section, overrides ...func(args mock.Arguments)) domain.SectionRepository {
		Db := new(mocks.Db)

		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Section)

			arg.SectionID = spy.SectionID
			arg.ModuleID = spy.ModuleID
			arg.Label = spy.Label
			arg.Lessons = spy.Lessons

			if len(overrides) > 0 {
				for _, override := range overrides {
					override(args)
				}
			}
		}

		var result domain.Section
		Db.On("Save", *spy, &result).Return().Run(mockResult)

		repo := repository.NewSectionRepository(Db)

		return repo
	}

	t.Run("should save a section", func(t *testing.T) {
		sectionSpy := entity_mocks.SectionMock()
		assert.Equal(t, sectionSpy.Label, "mock-label")
		sectionSpy.Label = "updated-label"

		repo := preSaveTest(sectionSpy)

		expected, err := repo.Save(*sectionSpy)

		assert.Nil(t, err)
		assert.NotNil(t, expected)
		assert.Equal(t, sectionSpy.SectionID, expected.SectionID)
		assert.Equal(t, "updated-label", expected.Label)
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
