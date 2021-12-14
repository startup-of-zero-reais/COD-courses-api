package domain

type (
	Db interface {
		Connect()
		Create(entity interface{}) interface{}
		Save(entity interface{}) interface{}
		Search(param map[string]string) []interface{}
		Delete(param map[string]string) bool
	}
)
