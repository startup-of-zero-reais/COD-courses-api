package repository

import (
	"errors"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/util"
)

type (
	LessonRepositoryImpl struct {
		Db domain.Db
	}
)

func NewLessonRepository(db domain.Db) *LessonRepositoryImpl {
	return &LessonRepositoryImpl{
		Db: db,
	}
}

func (l *LessonRepositoryImpl) Create(lesson domain.Lesson) (*domain.Lesson, error) {
	var result domain.Lesson
	l.Db.Create(lesson, &result)

	if result.LessonID == "" {
		return nil, errors.New("não foi possível criar aula na base de dados")
	}

	return &result, nil
}

func (l *LessonRepositoryImpl) Save(lesson domain.Lesson) (*domain.Lesson, error) {
	return nil, nil
}

func (l *LessonRepositoryImpl) Get(search map[string]string, pagination map[string]string) ([]domain.Lesson, error) {
	searchParam := util.MergeMaps(search, pagination)

	var result []domain.Lesson
	l.Db.Search(searchParam, &result)

	return result, nil
}

func (l *LessonRepositoryImpl) Delete(lessonID string) error {
	return nil
}
