package tests

import (
	"smalldoc124/service/internal/product"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/stretchr/testify/mock"
)

type mockDBProduct struct {
	mock.Mock
}

func (dbProduct *mockDBProduct) CreateNewProduct(db *sqlx.DB, newProduct product.NewProduct, now time.Time) (product.Product, error) {
	argument := dbProduct.Called(db, newProduct, now)
	return argument.Get(0).(product.Product), argument.Error(1)
}

func (dbProduct *mockDBProduct) GetProductByID(db *sqlx.DB, id string) (product.Product, error) {
	argument := dbProduct.Called(db, id)
	return argument.Get(0).(product.Product), argument.Error(1)
}
