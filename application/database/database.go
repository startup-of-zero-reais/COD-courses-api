package database

import (
	"fmt"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/util"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"log"
)
import "gorm.io/gorm"

type (
	Database struct {
		Db  *gorm.DB
		Env string
	}
)

func dsn(env string) string {
	if env == "testing" {
		return fmt.Sprintf("../test_db.sqlite")
	}

	user := util.GetEnv("MYSQL_USER", "root")
	pass := util.GetEnv("MYSQL_PASS", "root")
	host := util.GetEnv("MYSQL_HOST", "localhost")
	port := util.GetEnv("MYSQL_PORT", "3306")
	dbase := util.GetEnv("MYSQL_DATABASE", "courses")

	return fmt.Sprintf(
		"%s:%s@tpc(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbase,
	)
}

func NewDatabase() *Database {
	return &Database{
		Env: "development",
	}
}

func (d *Database) Connect() {
	if d.Env == "testing" {
		db, err := gorm.Open(sqlite.Open(dsn(d.Env)))
		if err != nil {
			log.Fatalf("erro ao iniciar conexao com banco: %s", err.Error())
		}

		d.Db = db
		return
	}

	db, err := gorm.Open(mysql.Open(dsn(d.Env)))
	if err != nil {
		log.Fatalf("erro ao iniciar conexao com banco: %s", err.Error())
	}

	d.Db = db
}

func (d *Database) Create(entity interface{}, result domain.Result) {
	d.Db.Create(entity).Scan(&result)
}

func (d *Database) Save(entity interface{}, result domain.Result) {
	d.Db.Save(entity).Scan(&result)
}

func (d *Database) Search(param map[string]string, result domain.Result) {
	d.Db.Find(&result, param).Scan(&result)
}

func (d *Database) Delete(param map[string]string, result domain.Result) bool {
	r := d.Db.Where(param).Delete(&result)
	return r.RowsAffected > 0
}
