package product

import "github.com/jmoiron/sqlx"

type ProductDB interface {
	GetProductByID(id string, db *sqlx.DB) (Product, error)
}
