package repository

import (
	"errors"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/util"
)

type (
	SectionRepositoryImpl struct {
		Db domain.Db
	}
)

func NewSectionRepository(Db domain.Db) domain.SectionRepository {
	return &SectionRepositoryImpl{Db: Db}
}

func (s *SectionRepositoryImpl) Create(section domain.Section) (*domain.Section, error) {
	var result domain.Section
	s.Db.Create(section, &result)

	if result.SectionID == "" {
		return nil, errors.New("erro ao criar seção")
	}

	return &result, nil
}

func (s *SectionRepositoryImpl) Save(section domain.Section) (*domain.Section, error) {
	if section.SectionID == "" {
		return nil, errors.New("erro ao salvar uma seção inexistente")
	}

	var result domain.Section
	s.Db.Save(section, &result)

	if result.SectionID == "" || &result == nil {
		return nil, errors.New("ocorreu algum erro ao salvar seção")
	}

	return &result, nil
}

func (s *SectionRepositoryImpl) Get(search map[string]string, pagination map[string]string) ([]domain.Section, error) {
	var result []domain.Section
	s.Db.Search(util.MergeMaps(search, pagination), &result)

	if len(result) <= 0 || &result == nil {
		return nil, errors.New("nenhuma seção encontrada")
	}

	return result, nil
}

func (s *SectionRepositoryImpl) Delete(sectionID string) error {
	var result domain.Section

	wasDeleted := s.Db.Delete(map[string]string{
		"section_id": sectionID,
	}, &result)

	if !wasDeleted {
		return errors.New("não foi possível deletar a seção")
	}

	return nil
}

func (s *SectionRepositoryImpl) Count() uint {
	return s.Db.TotalRows()
}
