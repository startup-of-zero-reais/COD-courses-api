package database

import (
	"fmt"
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	"github.com/startup-of-zero-reais/COD-courses-api/util"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"log"
	"strconv"
)
import "gorm.io/gorm"

type (
	Database struct {
		Db  *gorm.DB
		Dsn string
		Env string
	}
)

func (d *Database) BuildDsn() {
	if d.Env == "testing" {
		d.Dsn = fmt.Sprintf("./test_sqlite.db")
		return
	}

	user := util.GetEnv("MYSQL_USER", "root")
	pass := util.GetEnv("MYSQL_PASS", "root")
	host := util.GetEnv("MYSQL_HOST", "localhost")
	port := util.GetEnv("MYSQL_PORT", "3306")
	dbase := util.GetEnv("MYSQL_DATABASE", "courses")

	d.Dsn = fmt.Sprintf(
		"%s:%s@tpc(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbase,
	)
}

func NewDatabase(env ...string) *Database {
	dbase := &Database{Env: "development"}

	if len(env) > 0 {
		dbase.Env = env[0]
	}

	dbase.BuildDsn()
	return dbase
}

func (d *Database) Connect() {
	if d.Env == "testing" {
		db, err := gorm.Open(sqlite.Open(d.Dsn))
		if err != nil {
			fmt.Println(err.Error())
			log.Fatalf("erro ao iniciar conexao com banco de teste: %s", err.Error())
		}

		d.Db = db
		return
	}

	db, err := gorm.Open(mysql.Open(d.Dsn))
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
	offset := 0
	limit := 10

	for k, v := range param {
		switch k {
		case "per_page":
			perPage, _ := strconv.Atoi(v)
			limit = perPage
		case "page":
			page, _ := strconv.Atoi(v)
			if page <= 0 {
				page = 1
			}
			offset = (limit * page) - limit
		}
	}

	d.Db.Offset(offset).Limit(limit).Find(&result, param).Scan(&result)
}

func (d *Database) Delete(param map[string]string, result domain.Result) bool {
	r := d.Db.Where(param).Delete(&result)
	return r.RowsAffected > 0
}

func (d *Database) TotalRows() uint {
	var result interface{}
	r := d.Db.Select("1").FindInBatches(&result, 10000, func(tx *gorm.DB, batch int) error {
		return nil
	})

	return uint(r.RowsAffected)
}
