package mocks

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	"log"
)

type FakeArtifactDb struct {
	Db *sql.DB
}

func MockArtifact(artifactId, lessonId string) domain.Artifact {
	if artifactId == "" {
		artifactId = "artifact-mock-uuid"
	}

	if lessonId == "" {
		lessonId = "lesson-mock-uuid"
	}

	return domain.Artifact{
		ArtifactID: artifactId,
		LessonID:   lessonId,
		Link:       "mock-link",
	}
}

func (f *FakeArtifactDb) BeforeTest() {
	f.Connect()
	defer f.Db.Close()

	stmt, err := f.Db.Prepare(
		"CREATE TABLE IF NOT EXISTS artifacts (artifact_id varchar(36) primary key default((lower(hex(randomblob(4))) || '-' || lower(hex(randomblob(2))) || '-4' || substr(lower(hex(randomblob(2))),2) || '-' || substr('89ab',abs(random()) % 4 + 1, 1) || substr(lower(hex(randomblob(2))),2) || '-' || lower(hex(randomblob(6))))), lesson_id varchar(36), link varchar(255))",
	)
	if err != nil {
		log.Fatalf("erro ao preparar tabela: %s", err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("erro ao criar tabela: %s", err.Error())
	}
}

func (f *FakeArtifactDb) AfterSingleTest() {
	f.Connect()
	defer f.Db.Close()

	stmt, _ := f.Db.Prepare("DELETE FROM artifacts")
	_, err := stmt.Exec()
	if err != nil {
		log.Fatalf("erro ao limpar tabela: %s", err)
	}
}

func (f *FakeArtifactDb) Connect() {
	db, err := sql.Open("sqlite3", "../test_db.sqlite")
	if err != nil {
		log.Fatalf("erro ao conectar ao banco: %s", err.Error())
	}

	f.Db = db
}

func (f *FakeArtifactDb) Create(entity interface{}, result domain.Result) {
	artifact := entity.(domain.Artifact)

	f.Connect()
	defer f.Db.Close()

	stmt, err := f.Db.Prepare("INSERT INTO artifacts (lesson_id, link) VALUES (?, ?)")
	if err != nil {
		log.Fatalf("erro ao preparar query: %s", err)
	}

	_, err = stmt.Exec(
		artifact.LessonID, artifact.Link,
	)
	if err != nil {
		log.Fatalf("erro ao executar create: %s", err)
	}

	stmt, _ = f.Db.Prepare("SELECT * FROM artifacts WHERE lesson_id = ? AND link = ?")
	var modelResult domain.Artifact
	_ = stmt.QueryRow(artifact.LessonID, artifact.Link).Scan(
		&modelResult.ArtifactID, &modelResult.LessonID, &modelResult.Link,
	)

	result = modelResult
}

func (f *FakeArtifactDb) Save(entity interface{}, result domain.Result) {
	var artifacts []domain.Artifact

	f.Search(map[string]string{
		"artifact_id": (entity.(domain.Artifact)).ArtifactID,
	}, &artifacts)

	f.Connect()
	defer f.Db.Close()

	if len(artifacts) > 0 {
		stmt, err := f.Db.Prepare("UPDATE artifacts SET lesson_id=?, link=? WHERE artifact_id=?")
		if err != nil {
			log.Fatalf("erro ao preparar query: %s", err)
		}

		a := entity.(domain.Artifact)
		_, err = stmt.Exec(a.LessonID, a.Link, a.ArtifactID)
		if err != nil {
			log.Fatalf("erro ao executar a query: %s", err.Error())
		}

		var modelResult domain.Artifact
		stmt, _ = f.Db.Prepare("SELECT * FROM artifacts WHERE artifact_id = ?")
		_ = stmt.QueryRow(a.ArtifactID).Scan(
			&modelResult.ArtifactID, &modelResult.LessonID, &modelResult.Link,
		)

		result = modelResult
		return
	}

	f.Create(entity, result)

	return
}

func (f *FakeArtifactDb) Search(param map[string]string, result domain.Result) {
	var artifacts []interface{}

	f.Connect()
	defer f.Db.Close()

	for k, p := range param {
		query := fmt.Sprintf("SELECT * FROM artifacts WHERE %s=?", k)
		stmt, err := f.Db.Prepare(query)
		if err != nil {
			log.Fatalf("erro ao preparar statement: %s", err.Error())
		}

		var e domain.Artifact
		_ = stmt.QueryRow(p).Scan(&e.ArtifactID, &e.LessonID, &e.Link)

		if e.ArtifactID != "" {
			artifacts = append(artifacts, e)
		}
	}

	result = artifacts

	return
}

func (f *FakeArtifactDb) Delete(param map[string]string, result domain.Result) bool {
	wasDeleted := false

	var isSet []struct {
	}
	f.Search(param, &isSet)
	f.Connect()
	defer f.Db.Close()

	for k, p := range param {
		if len(isSet) > 0 {
			query := fmt.Sprintf("DELETE FROM artifacts WHERE %s=?", k)
			stmt, err := f.Db.Prepare(query)
			if err != nil {
				log.Fatalf("erro ao preparar o delete: %s", err.Error())
			}

			result, err := stmt.Exec(p)
			if err != nil {
				log.Fatalf("ocorreu um erro ao executar a query: %s", err.Error())
			}

			if rows, _ := result.RowsAffected(); rows > 0 {
				wasDeleted = true
			}
		}
	}

	return wasDeleted
}
