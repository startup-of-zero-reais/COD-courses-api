package database

import "database/sql"

type (
	Database struct {
	}
)

func (d *Database) Connect() *sql.DB {
	return nil
}

func (d *Database) Create(entity interface{}) interface{} {
	return nil
}

func (d *Database) Save(entity interface{}) interface{} {
	return nil
}

func (d *Database) Search(param map[string]string) []interface{} {
	return nil
}

func (d *Database) Delete(param map[string]string) bool {
	return false
}
