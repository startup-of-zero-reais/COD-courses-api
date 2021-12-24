package entity_mocks

import (
	"github.com/google/uuid"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	"time"
)

func LessonMock(overrideCols ...map[string]interface{}) *domain.Lesson {
	lesson := &domain.Lesson{
		LessonID:      "",
		SectionID:     uuid.NewString(),
		VideoSource:   "https://video.source/mock",
		DurationTotal: 150,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if len(overrideCols) > 0 {
		toOverride := overrideCols[0]
		for key, value := range toOverride {
			switch key {
			case "lesson_id":
				lesson.LessonID = value.(string)
			case "section_id":
				lesson.SectionID = value.(string)
			case "video_source":
				lesson.VideoSource = value.(string)
			case "duration_total":
				lesson.DurationTotal = value.(uint)
			case "created_at":
				lesson.CreatedAt = value.(time.Time)
			case "updated_at":
				lesson.UpdatedAt = value.(time.Time)
			}
		}
	}

	return lesson
}
