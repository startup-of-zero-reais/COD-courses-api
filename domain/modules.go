package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Module struct {
		ModuleID  string    `json:"module_id,omitempty" paginator:"key:module_id" gorm:"type:varchar(36);primaryKey;column:module_id;"`
		CourseID  string    `json:"course_id,omitempty" paginator:"skey:course_id" gorm:"type:varchar(36);column:course_id;"`
		Sections  []Section `json:"sections,omitempty"`
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`

		Self     string `json:"_self" paginator:"_self" gorm:"-"`
		Embedded string `json:"_embedded" paginator:"_embedded" gorm:"-"`

		SectionOrders []string `json:"section_orders,omitempty" gorm:"column:section_orders"`
	}

	ModuleRepository interface {
		Create(module Module) (*Module, error)
		Save(module Module) (*Module, error)
		Get(search map[string]string, pagination map[string]string) ([]Module, error)
		Delete(moduleID string) error
		Counter
	}

	ModuleService interface {
		Add(module Module) (*Module, error)
		Save(module Module) (*Module, error)
		ListByCourses(courseID string, pagination map[string]string) ([]Module, error)
		Get(moduleID string) (*Module, error)
		Delete(moduleID string) error
	}
)

func (m *Module) BeforeCreate(tx *gorm.DB) (err error) {
	m.ModuleID = uuid.NewString()
	return
}
