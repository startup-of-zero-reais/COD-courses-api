package repository

import (
	"errors"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
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
	
	return &result, nil
}

func (s *SectionRepositoryImpl) Get(search map[string]string, pagination map[string]string) ([]domain.Section, error) {
	return nil, nil
}

func (s *SectionRepositoryImpl) Delete(sectionID string) error {
	return nil
}

func (s *SectionRepositoryImpl) Count() uint {
	return uint(0)
}
