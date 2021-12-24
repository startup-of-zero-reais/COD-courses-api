package repository

import (
	"errors"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
)

type ArtifactRepositoryImpl struct {
	Db domain.Db
}

func NewArtifactRepository(db domain.Db) *ArtifactRepositoryImpl {
	return &ArtifactRepositoryImpl{
		Db: db,
	}
}

func (a *ArtifactRepositoryImpl) Create(artifact domain.Artifact) (*domain.Artifact, error) {
	var result domain.Artifact
	a.Db.Create(artifact, &result)

	if result.ArtifactID == "" || &result == nil {
		return nil, errors.New("falha ao criar artefato na base de dados")
	}

	return &result, nil
}

func (a *ArtifactRepositoryImpl) Save(artifact domain.Artifact) (*domain.Artifact, error) {
	if artifact.ArtifactID == "" {
		return nil, errors.New("artefato sem referência no banco")
	}

	var artifactSaved domain.Artifact
	a.Db.Save(artifact, &artifactSaved)
	if &artifactSaved == nil || artifactSaved.ArtifactID == "" {
		return nil, errors.New("ocorreu algum erro ao salvar artefato")
	}

	return &artifactSaved, nil
}

func (a *ArtifactRepositoryImpl) Get(searchParam map[string]string, pagination map[string]string) ([]domain.Artifact, error) {
	var artifacts []domain.Artifact

	searchParams := map[string]string{}
	for key, value := range searchParam {
		searchParams[key] = value
	}
	for key, value := range pagination {
		searchParams[key] = value
	}

	a.Db.Search(searchParams, &artifacts)

	return artifacts, nil
}

func (a *ArtifactRepositoryImpl) Delete(artifactId string) error {
	wasDeleted := a.Db.Delete(map[string]string{"artifact_id": artifactId}, &domain.Artifact{})

	if !wasDeleted {
		return errors.New("ocorreu um erro ao deletar artefato da base de dados")
	}

	return nil
}
