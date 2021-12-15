package mocks

import (
	"github.com/startup-of-zero-reais/COD-courses-api/application/database"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	"gorm.io/gorm"
)

type (
	DatabaseUtil struct {
		Db *gorm.DB
		*database.Database
	}
)

func SetupTest() *DatabaseUtil {
	db := database.NewDatabase()
	db.Env = "testing"
	db.Connect()

	dbUtil := &DatabaseUtil{}
	dbUtil.Database = db
	dbUtil.Db = dbUtil.Database.Db

	return dbUtil
}

func (d *DatabaseUtil) BeforeTests() {
	d.ClearDB()
}

func (d *DatabaseUtil) AfterTests() {
	defer d.ClearDB()
}

func (d *DatabaseUtil) ClearDB() {
	d.Db.Where("1 = ?", 1).Delete(&domain.Artifact{})
}
