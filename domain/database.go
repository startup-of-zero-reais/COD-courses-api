package domain

type (
	Result interface{}

	Db interface {
		Connect()
		Create(entity interface{}, result Result)
		Save(entity interface{}, result Result)
		Search(param map[string]string, result Result)
		Delete(param map[string]string, result Result) bool
		TotalRows() uint
	}
)
