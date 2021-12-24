package service

import (
	"errors"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
)

type (
	LessonServiceImpl struct {
		domain.LessonRepository
	}
)

func NewLessonService(repo domain.LessonRepository) *LessonServiceImpl {
	return &LessonServiceImpl{
		LessonRepository: repo,
	}
}

func (l *LessonServiceImpl) Add(lesson domain.Lesson) (*domain.Lesson, error) {
	return nil, nil
}

func (l *LessonServiceImpl) Save(lesson domain.Lesson) (*domain.Lesson, error) {
	result, err := l.LessonRepository.Save(lesson)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (l *LessonServiceImpl) ListBySection(sectionID string, query map[string]string) ([]domain.Lesson, error) {
	search := map[string]string{
		"section_id": sectionID,
	}

	result, err := l.LessonRepository.Get(search, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (l *LessonServiceImpl) Get(lessonID string) (*domain.Lesson, error) {
	result, err := l.LessonRepository.Get(map[string]string{
		"lesson_id": lessonID,
	}, map[string]string{"page": "1", "per_page": "1"})

	if err != nil {
		return nil, err
	}

	if len(result) <= 0 {
		return nil, errors.New("nenhuma aula encontrada")
	}

	if len(result) > 1 {
		return nil, errors.New("mais aulas encontradas. contate o administrador do sistema")
	}

	lesson := result[0]
	return &lesson, nil
}

func (l *LessonServiceImpl) Delete(lessonID string) error {
	return l.LessonRepository.Delete(lessonID)
}
