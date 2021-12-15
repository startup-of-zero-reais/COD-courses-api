package repository

import (
	"errors"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
)

type ArtifactRepositoryImpl struct {
	Db domain.Db
}

func (a *ArtifactRepositoryImpl) Create(artifact domain.Artifact) (*domain.Artifact, error) {
	var result domain.Artifact
	a.Db.Create(artifact, &result)
	return &result, nil
}

func (a *ArtifactRepositoryImpl) Save(artifact domain.Artifact) (*domain.Artifact, error) {
	var artifactSaved domain.Artifact
	a.Db.Save(artifact, &artifactSaved)
	if &artifactSaved == nil || artifactSaved.ArtifactID == "" {
		return nil, errors.New("ocorreu algum erro ao salvar artefato")
	}

	return &artifactSaved, nil
}

func (a *ArtifactRepositoryImpl) Get(searchParam map[string]string) ([]domain.Artifact, error) {
	var artifacts []domain.Artifact
	a.Db.Search(searchParam, &artifacts)

	return artifacts, nil
}

func (a *ArtifactRepositoryImpl) Delete(artifactId string) error {
	wasDeleted := a.Db.Delete(map[string]string{"artifact_id": artifactId}, domain.Artifact{})

	if !wasDeleted {
		return errors.New("ocorreu um erro ao deletar artefato da base de dados")
	}

	return nil
}
