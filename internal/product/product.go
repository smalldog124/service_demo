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

	const query = `INSERT INTO products (product_id,name, price, amount, date_created, date_updated)VALUES ($1, $2, $3, $4, $5, $6)`
	tx := db.MustBegin()
	tx.MustExec(query, product.ID, product.Name, product.Price, product.Amount, product.DateCreated, product.DateUpdated)
	if err := tx.Commit(); err != nil {
		return Product{}, err
	}
	return product, nil
}

func ListProduct(db *sqlx.DB) ([]Product, error) {
	var product []Product
	const query = `SELECT product_id,name, price, amount, date_created, date_updated FROM products`
	err := db.Select(&product, query)
	if err != nil {
		return []Product{}, err
	}
	for index, prod := range product {
		product[index].DateCreated = prod.DateCreated.UTC()
		product[index].DateUpdated = prod.DateUpdated.UTC()
	}
	return product, nil
}

func GetProductByID(db *sqlx.DB, id string) (Product, error) {
	var product Product
	const query = `SELECT product_id,name, price, amount, date_created, date_updated FROM products WHERE product_id=$1`
	err := db.Get(&product, query, id)
	if err != nil {
		return Product{}, err
	}
	product.DateCreated = product.DateCreated.UTC()
	product.DateUpdated = product.DateUpdated.UTC()
	return product, nil
}

func Update(db *sqlx.DB, id string, update UpdateProduct, now time.Time) error {
	product, err := GetProductByID(db, id)
	if err != nil {
		return err
	}

	if update.Name != nil {
		product.Name = *update.Name
	}
	if update.Price != nil {
		product.Price = *update.Price
	}
	if update.Amount != nil {
		product.Amount = *update.Amount
	}
	product.DateUpdated = now
	const query = `UPDATE products SET "name" = $2, "price" = $3, "amount" = $4, "date_updated" = $5 WHERE product_id=$1`
	tx := db.MustBegin()
	tx.MustExec(query, product.ID, product.Name, product.Price, product.Amount, product.DateUpdated)
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
