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
	if lesson.LessonID == "" {
		return nil, errors.New("esta aula nao possui registro na base de dados")
	}

	var result domain.Lesson
	l.Db.Save(lesson, &result)

	if result.LessonID == "" || result.LessonID != lesson.LessonID {
		return nil, errors.New("nao foi possível atualizar o registro")
	}

	return &result, nil
}

func (l *LessonRepositoryImpl) Get(search map[string]string, pagination map[string]string) ([]domain.Lesson, error) {
	searchParam := util.MergeMaps(search, pagination)

	var result []domain.Lesson
	l.Db.Search(searchParam, &result)

	return result, nil
}

func (l *LessonRepositoryImpl) Delete(lessonID string) error {
	var result domain.Lesson
	wasDeleted := l.Db.Delete(map[string]string{
		"lesson_id": lessonID,
	}, &result)

	if !wasDeleted {
		return errors.New("erro ao deletar aula")
	}

	return nil
}
