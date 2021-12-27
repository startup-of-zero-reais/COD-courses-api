package service

import "github.com/startup-of-zero-reais/COD-courses-api/domain"

type SectionServiceImpl struct {
	domain.SectionRepository
}

func NewSectionService(repo domain.SectionRepository) domain.SectionService {
	return &SectionServiceImpl{
		SectionRepository: repo,
	}
}

func (s *SectionServiceImpl) Add(section domain.Section) (*domain.Section, error) {
	return s.SectionRepository.Create(section)
}

func (s *SectionServiceImpl) Save(section domain.Section) (*domain.Section, error) {
	return nil, nil
}

func (s *SectionServiceImpl) ListByModule(moduleID string, pagination map[string]string) ([]domain.Section, error) {
	return nil, nil
}

func (s *SectionServiceImpl) Get(sectionID string) (*domain.Section, error) {
	return nil, nil
}

func (s *SectionServiceImpl) Delete(sectionID string) error {
	return nil
}