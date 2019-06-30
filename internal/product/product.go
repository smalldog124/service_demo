package product

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type ProductDB interface {
	GetProductByID(id string, db *sqlx.DB) (Product, error)
}

func CreateNewProduct(db *sqlx.DB, newProduct NewProduct, now time.Time) (Product, error) {
	product := Product{
		ID:          "1",
		Name:        newProduct.Name,
		Price:       newProduct.Price,
		Amount:      newProduct.Amount,
		DateCreated: now.UTC(),
		DateUpdated: now.UTC(),
	}

	const query = `INSERT INTO products (id,name, price, amount, date_created, date_updated)VALUES ($1, $2, $3, $4, $5, $6)`
	tx := db.MustBegin()
	tx.MustExec(query, product.ID, product.Name, product.Price, product.Amount, product.DateCreated, product.DateUpdated)
	if err := tx.Commit(); err != nil {
		return Product{}, err
	}
	return product, nil
}
