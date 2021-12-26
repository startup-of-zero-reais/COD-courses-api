package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Section struct {
		SectionID string `json:"section_id,omitempty" paginator:"key:section_id" gorm:"type:varchar(36);primaryKey;column:section_id;"`
		ModuleID  string `json:"module_id,omitempty" paginator:"skey:module_id" gorm:"type:varchar(36);column:module_id;"`
		Label     string `json:"label,omitempty" paginator:"-" gorm:"column:label"`

		Self     string `json:"_self,omitempty" paginator:"_self" gorm:"-"`
		Embedded string `json:"_embedded,omitempty" paginator:"_embedded" gorm:"-"`

		Lessons []Lesson `json:"lessons,omitempty"`
	}

	SectionRepository interface {
		Create(section Section) (*Section, error)
		Save(section Section) (*Section, error)
		Get(search map[string]string, pagination map[string]string) ([]Section, error)
		Delete(sectionID string) error
		Counter
	}

	SectionService interface {
		Add(section Section) (*Section, error)
		Save(section Section) (*Section, error)
		ListByModule(moduleID string, pagination map[string]string) ([]Section, error)
		Get(sectionID string) (*Section, error)
		Delete(sectionID string) error
	}
)

func (s *Section) BeforeCreate(tx *gorm.DB) (err error) {
	s.SectionID = uuid.NewString()
	return
}
