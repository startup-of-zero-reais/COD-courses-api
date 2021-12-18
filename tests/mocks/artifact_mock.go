package mocks

import "github.com/startup-of-zero-reais/COD-courses-api/domain"

func SetNewIfEmpty(_default, _new string) string {
	if _new != "" {
		_default = _new
	}

	return _default
}

func ArtifactMock(args ...string) *domain.Artifact {
	artifactID := ""
	lessonID := ""
	link := ""
	if len(args) > 0 {
		artifactID = args[0]
	}

	if len(args) > 1 {
		lessonID = args[1]
	}

	if len(args) > 2 {
		link = args[2]
	}

	return &domain.Artifact{
		ArtifactID: SetNewIfEmpty("", artifactID),
		LessonID:   SetNewIfEmpty("mock-lesson-id", lessonID),
		Link:       SetNewIfEmpty("mock-link", link),
	}
}
