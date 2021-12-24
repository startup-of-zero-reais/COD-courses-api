package service

import "github.com/startup-of-zero-reais/COD-courses-api/domain"

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
	return nil, nil
}

func (l *LessonServiceImpl) ListBySection(sectionID string, query map[string]string) ([]domain.Lesson, error) {
	return nil, nil
}

func (l *LessonServiceImpl) Get(lessonID string) (*domain.Lesson, error) {
	return nil, nil
}

func (l *LessonServiceImpl) Delete(lessonID string) error {
	return l.LessonRepository.Delete(lessonID)
}
