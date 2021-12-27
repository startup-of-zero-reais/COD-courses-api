package service_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/startup-of-zero-reais/COD-courses-api/application/service"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	mocks "github.com/startup-of-zero-reais/COD-courses-api/mocks/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/tests/entity_mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type TestSuit struct {
	Meth          string
	Repo          *mocks.SectionRepository
	Args          []interface{}
	MockExecution []func(args mock.Arguments)
}

func BeforeAllSection(meth string) *TestSuit {
	return &TestSuit{
		Meth: meth,
		Repo: new(mocks.SectionRepository),
	}
}

func (t *TestSuit) BeforeEach(args ...interface{}) {
	t.Args = args
}

func (t *TestSuit) AfterEach() {
	t.Args = nil
}

func (t *TestSuit) MockReturns(returns ...interface{}) {
	t.Repo.On(t.Meth, t.Args...).Return(returns...).Run(func(args mock.Arguments) {
		if len(t.MockExecution) > 0 {
			for _, override := range t.MockExecution {
				override(args)
			}
		}
	})
}

func (t *TestSuit) PutExecutionMock(m func(args mock.Arguments)) *TestSuit {
	t.MockExecution = append(t.MockExecution, m)
	return t
}

func (t *TestSuit) getService() domain.SectionService {
	return service.NewSectionService(t.Repo)
}

func TestSectionServiceImpl_Add(t *testing.T) {
	ts := BeforeAllSection("Create")

	t.Run("should add a section", func(t *testing.T) {
		sectionSpy := *entity_mocks.SectionMock(map[string]interface{}{
			"section_id": "-",
			"lessons":    []domain.Lesson{},
		})
		assert.Zero(t, sectionSpy.SectionID)
		sectionSpy.SectionID = uuid.NewString()

		ts.BeforeEach(sectionSpy)
		defer ts.AfterEach()

		ts.MockReturns(&sectionSpy, nil)
		svc := ts.getService()

		expected, err := svc.Add(sectionSpy)

		assert.Nil(t, err)
		assert.NotNil(t, expected)
		assert.NotZero(t, expected.SectionID)
	})
	t.Run("should fail on add a section", func(t *testing.T) {
		sectionSpy := *entity_mocks.SectionMock()

		ts.BeforeEach(sectionSpy)
		defer ts.AfterEach()

		ts.MockReturns(nil, errors.New("erro ao criar seção"))
		svc := ts.getService()

		expected, err := svc.Add(sectionSpy)

		assert.Nil(t, expected)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "erro ao criar seção")
	})
}

func TestSectionServiceImpl_Get(t *testing.T) {
	ts := BeforeAllSection("Get")

	t.Run("should get single section", func(t *testing.T) {
		sectionSpy := []domain.Section{
			*entity_mocks.SectionMock(map[string]interface{}{
				"section_id": "section-id",
			}),
		}
		search := map[string]string{
			"section_id": "section-id",
		}
		pagination := map[string]string{
			"page":     "1",
			"per_page": "1",
		}

		ts.BeforeEach(search, pagination)
		defer ts.AfterEach()

		ts.MockReturns(sectionSpy, nil)

		svc := ts.getService()

		expected, err := svc.Get("section-id")

		assert.Nil(t, err)
		assert.NotNil(t, expected)
		assert.Equal(t, expected.SectionID, "section-id")
	})
	t.Run("should fail if has no section", func(t *testing.T) {

	})
	t.Run("should fail if has more than one section", func(t *testing.T) {

	})
}

func TestSectionServiceImpl_ListByModule(t *testing.T) {
	t.Run("should list sections", func(t *testing.T) {

	})
	t.Run("should return empty slice if has no sections", func(t *testing.T) {

	})
}

func TestSectionServiceImpl_Save(t *testing.T) {
	t.Run("should update a section", func(t *testing.T) {

	})
	t.Run("should fail if does not exists the section was updating", func(t *testing.T) {

	})
}

func TestSectionServiceImpl_Delete(t *testing.T) {
	t.Run("should delete a section", func(t *testing.T) {

	})
	t.Run("should fail if no delete", func(t *testing.T) {

	})
}
