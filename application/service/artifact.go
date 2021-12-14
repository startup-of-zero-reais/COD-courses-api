package service

import (
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
)

type ArtifactServiceImpl struct {
	Repo domain.ArtifactRepository
}

func (s *ArtifactServiceImpl) Add(artifact domain.Artifact) (*domain.Artifact, error) {
	newArtifact, err := s.Repo.Create(artifact)

	if err != nil {
		return nil, err
	}

	return newArtifact, nil
}

func (s *ArtifactServiceImpl) List(lessonId string) ([]domain.Artifact, error) {
	var artifacts []domain.Artifact

	artifacts, err := s.Repo.Get(map[string]string{"lesson_id": lessonId})
	if err != nil {
		return nil, err
	}

	return artifacts, nil
}

func (s *ArtifactServiceImpl) Remove(artifactId string) error {
	return s.Repo.Delete(artifactId)
}
