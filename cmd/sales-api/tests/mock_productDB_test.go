package tests

import (
	"smalldoc124/service/internal/product"
	"time"

	"github.com/stretchr/testify/mock"
)

type mockDBProduct struct {
	mock.Mock
}

func (dbProduct *mockDBProduct) CreateNewProduct(newProduct product.NewProduct, now time.Time) (product.Product, error) {
	argument := dbProduct.Called(newProduct, now)
	return argument.Get(0).(product.Product), argument.Error(1)
}

func (dbProduct *mockDBProduct) GetProductByID(id string) (product.Product, error) {
	argument := dbProduct.Called(id)
	return argument.Get(0).(product.Product), argument.Error(1)
}
