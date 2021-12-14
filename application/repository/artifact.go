package repository

import (
	"errors"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
)

type ArtifactRepositoryImpl struct {
	Db domain.Db
}

func (a *ArtifactRepositoryImpl) Create(artifact domain.Artifact) (*domain.Artifact, error) {
	result := a.Db.Create(artifact)
	artifactResult := result.(domain.Artifact)
	return &artifactResult, nil
}

func (a *ArtifactRepositoryImpl) Save(artifact domain.Artifact) (*domain.Artifact, error) {
	artifactSaved := a.Db.Save(artifact)
	if artifactSaved == nil {
		return nil, errors.New("ocorreu algum erro ao salvar artefato")
	}

	pointer := artifactSaved.(domain.Artifact)

	return &pointer, nil
}

func (a *ArtifactRepositoryImpl) Get(searchParam map[string]string) ([]domain.Artifact, error) {
	artifactsInterface := a.Db.Search(searchParam)
	var artifacts []domain.Artifact

	for _, ai := range artifactsInterface {
		artifacts = append(artifacts, ai.(domain.Artifact))
	}

	return artifacts, nil
}

func (a *ArtifactRepositoryImpl) Delete(artifactId string) error {
	wasDeleted := a.Db.Delete(map[string]string{"artifact_id": artifactId})

	if !wasDeleted {
		return errors.New("ocorreu um erro ao deletar artefato da base de dados")
	}

	return nil
}
