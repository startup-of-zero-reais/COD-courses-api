package domain

type (
	Artifact struct {
		ArtifactID string
		LessonID   string
		Link       string
	}

	ArtifactRepository interface {
		Create(artifact Artifact) (*Artifact, error)
		Save(artifact Artifact) (*Artifact, error)
		Get(searchParam map[string]string) ([]Artifact, error)
		Delete(artifactId string) error
	}

	ArtifactService interface {
		Add(artifact Artifact) (*Artifact, error)
		List(lessonId string) ([]Artifact, error)
		Remove(artifactId string) error
	}
)
