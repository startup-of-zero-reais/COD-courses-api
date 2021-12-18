package database_test

import (
	"github.com/startup-of-zero-reais/COD-courses-api/application/database"
	"github.com/stretchr/testify/require"
	"log"
	"os/exec"
	"testing"
)

func AfterTests() {
	cmd := exec.Command("rm", "-rf", "./test_sqlite.db")
	err := cmd.Run()
	if err != nil {
		log.Printf("erro ao rm -rf: %s", err.Error())
	}
}

func TestDatabase_BuildDsn(t *testing.T) {
	t.Run("should build dsn on testing", func(t *testing.T) {
		Db := database.NewDatabase("testing")
		Db.BuildDsn()

		require.Equal(t, Db.Dsn, "./test_sqlite.db")
	})
	t.Run("should build dsn on prod/dev", func(t *testing.T) {
		Db := database.NewDatabase()
		Db.BuildDsn()

		require.Equal(t, Db.Dsn, "root:root@tpc(localhost:3306)/courses?charset=utf8mb4&parseTime=True&loc=Local")
	})
}

func TestDatabase_Connect(t *testing.T) {
	t.Run("should call connect method", func(t *testing.T) {
		defer AfterTests()

		Db := database.NewDatabase("testing")
		Db.Connect()

		require.NotNil(t, Db.Db)
	})
}
