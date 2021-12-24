package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Artifact struct {
		ArtifactID string `json:"artifact_id" paginator:"key:artifact_id" gorm:"type:varchar(36);primaryKey;column:artifact_id;"`
		LessonID   string `json:"lesson_id" gorm:"type:varchar(36);column:lesson_id"`
		Link       string `json:"_link" paginator:"_self" gorm:"type:varchar(255);column:link"`
	}

	ArtifactRepository interface {
		Create(artifact Artifact) (*Artifact, error)
		Save(artifact Artifact) (*Artifact, error)
		Get(searchParam map[string]string, pagination map[string]string) ([]Artifact, error)
		Delete(artifactId string) error
	}

	ArtifactService interface {
		Add(artifact Artifact) (*Artifact, error)
		List(lessonId string, query map[string]string) ([]Artifact, error)
		Remove(artifactId string) error
	}
)

func (a *Artifact) BeforeCreate(tx *gorm.DB) (err error) {
	a.ArtifactID = uuid.New().String()
	return
}
