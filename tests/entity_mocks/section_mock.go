package entity_mocks

import (
	"github.com/google/uuid"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
)

func SectionMock(overrides ...map[string]interface{}) *domain.Section {
	section := &domain.Section{
		SectionID: uuid.NewString(),
		ModuleID:  uuid.NewString(),
		Label:     "mock-label",
		Lessons: []domain.Lesson{
			*LessonMock(map[string]interface{}{"lesson_id": uuid.NewString()}),
			*LessonMock(map[string]interface{}{"lesson_id": uuid.NewString()}),
		},
	}

	if len(overrides) > 0 {
		override := overrides[0]
		for key, value := range override {
			switch key {
			case "section_id":
				if key == "-" {
					section.SectionID = ""
				} else {
					section.SectionID = value.(string)
				}
			case "module_id":
				section.ModuleID = value.(string)
			case "lessons":
				section.Lessons = value.([]domain.Lesson)
			case "label":
				section.Label = value.(string)
			}
		}
	}

	return section
}
