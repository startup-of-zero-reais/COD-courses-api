package mocks

import (
	"github.com/startup-of-zero-reais/COD-courses-api/domain"
	"github.com/stretchr/testify/mock"
)

type DbMock struct {
	mock.Mock
}

func (d *DbMock) Connect() {
	d.Called()
	return
}

func (d *DbMock) Create(entity interface{}, result domain.Result) {
	d.Called(entity, result)
	return
}

func (d *DbMock) Save(entity interface{}, result domain.Result) {
	d.Called(entity, result)
	return
}

func (d *DbMock) Search(param map[string]string, result domain.Result) {
	d.Called(param, result)
	return
}

func (d *DbMock) Delete(param map[string]string, result domain.Result) bool {
	args := d.Called(param, result)
	return args.Bool(0)
}
