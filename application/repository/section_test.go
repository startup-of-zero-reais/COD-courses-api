package repository_test

import (
	"github.com/google/uuid"
	"github.com/startup-of-zero-reais/COD-courses-api/application/repository"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	mocks "github.com/startup-of-zero-reais/COD-courses-api/mocks/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/tests/entity_mocks"
	"github.com/startup-of-zero-reais/COD-courses-api/util"
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
	t.Run("should fail on save if SectionID is empty", func(t *testing.T) {
		sectionSpy := entity_mocks.SectionMock(map[string]interface{}{
			// Deixa vazio a section_id
			"section_id": "-",
		})

		repo := preSaveTest(sectionSpy)

		expected, err := repo.Save(*sectionSpy)

		assert.Nil(t, expected)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "erro ao salvar uma seção inexistente")
	})
	t.Run("should fail if cannot save section", func(t *testing.T) {
		sectionSpy := entity_mocks.SectionMock()
		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Section)
			arg.SectionID = ""
			arg = &domain.Section{}
		}
		repo := preSaveTest(sectionSpy, mockResult)

		expected, err := repo.Save(*sectionSpy)
		assert.Nil(t, expected)
		assert.EqualError(t, err, "ocorreu algum erro ao salvar seção")
	})
}

func TestSectionRepositoryImpl_Get(t *testing.T) {
	preGetTest := func(sectionID string, results []domain.Section) (domain.SectionRepository, map[string]string) {
		Db := new(mocks.Db)
		search := map[string]string{
			"section_id": sectionID,
		}
		pagination := map[string]string{
			"page":     "1",
			"per_page": "10",
		}

		mockResult := func(args mock.Arguments) {
			arg := args.Get(1).(*[]domain.Section)

			_ = util.ReflectSlice(arg, results)
		}

		var res []domain.Section
		Db.On("Search", util.MergeMaps(search, pagination), &res).Return().Run(mockResult)

		return repository.NewSectionRepository(Db), pagination
	}

	t.Run("should return a single section", func(t *testing.T) {
		repo, pagination := preGetTest("section-id", []domain.Section{
			*entity_mocks.SectionMock(),
		})
		search := map[string]string{
			"section_id": "section-id",
		}

		expected, err := repo.Get(search, pagination)

		assert.Nil(t, err)
		assert.NotNil(t, expected)
		assert.Len(t, expected, 1)
	})
	t.Run("should fail if has no section", func(t *testing.T) {
		repo, pagination := preGetTest("not-exists-section", []domain.Section{})
		search := map[string]string{
			"section_id": "not-exists-section",
		}
		expected, err := repo.Get(search, pagination)

		assert.Nil(t, expected)
		assert.EqualError(t, err, "nenhuma seção encontrada")
	})
}

func TestSectionRepositoryImpl_Delete(t *testing.T) {
	preDeleteTest := func(sectionID string, returns bool) domain.SectionRepository {
		Db := new(mocks.Db)

		Db.On("Delete", map[string]string{
			"section_id": sectionID,
		}, &domain.Section{}).Return(returns)

		return repository.NewSectionRepository(Db)
	}
	t.Run("should delete a section", func(t *testing.T) {
		repo := preDeleteTest("section-id", true)

		err := repo.Delete("section-id")
		assert.Nil(t, err)
	})
	t.Run("should fail on try delete", func(t *testing.T) {
		repo := preDeleteTest("section-id", false)

		err := repo.Delete("section-id")

		assert.NotNil(t, err)
		assert.EqualError(t, err, "não foi possível deletar a seção")
	})
}

func TestSectionRepositoryImpl_Count(t *testing.T) {

}
