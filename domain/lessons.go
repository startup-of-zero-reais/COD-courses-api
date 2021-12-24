package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Lesson struct {
		LessonID      string `json:"lesson_id,omitempty" paginator:"key:lesson_id" gorm:"type:varchar(36);primaryKey;column:lesson_id;"`
		SectionID     string `json:"section_id,omitempty" paginator:"skey:section_id" gorm:"type:varchar(36);column:section_id"`
		VideoSource   string `json:"video_source,omitempty" gorm:"column:video_source"`
		DurationTotal uint   `json:"duration_total,omitempty" gorm:"column:duration_total"`

		Self     string `json:"_self,omitempty" paginator:"_self"`
		Embedded string `json:"_embedded,omitempty" paginator:"_embedded"`

		Artifacts []Artifact `json:"artifacts,omitempty"`

		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	LessonRepository interface {
		Create(lesson Lesson) (*Lesson, error)
		Save(lesson Lesson) (*Lesson, error)
		Get(search map[string]string, pagination map[string]string) ([]Lesson, error)
		Delete(lessonID string) error
		Counter
	}

	LessonService interface {
		Add(lesson Lesson) (*Lesson, error)
		Update(lesson Lesson) (*Lesson, error)
		ListBySection(sectionID string, query map[string]string) ([]Lesson, error)
		Get(lessonID string) (*Lesson, error)
		Delete(lessonID string) error
	}
)

func (l *Lesson) BeforeCreate(tx *gorm.DB) (err error) {
	l.LessonID = uuid.NewString()
	return
}
