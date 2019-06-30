package tests

import (
	"log"
	"smalldoc124/service/internal/database"
	"smalldoc124/service/internal/product"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ProductDB(t *testing.T) {
	cfg := database.Config{
		User:       "smalldoc",
		Password:   "example",
		Host:       "localhost",
		Port:       5432,
		Name:       "smalldoc",
		DisableTLS: true,
	}
	db, err := database.Open(cfg)
	if err != nil {
		log.Fatal("connect db error: ", err)
	}
	defer db.Close()

	schema := `
	CREATE TABLE products (
		product_id   TEXT,
		name         TEXT,
		price         DECIMAL,
		amount     INT,
		date_created TIMESTAMP,
		date_updated TIMESTAMP,
	
		PRIMARY KEY (product_id)
	);
	`
	resual, err := database.NewTable(db, schema)
	if err != nil {
		t.Log("ceate tabel error: ", err)

	}
	t.Log(resual)
	t.Run("Create New Product", func(t *testing.T) {
		now := time.Date(2019, time.June, 1, 0, 0, 0, 0, time.UTC)
		expeted := product.Product{
			ID:          "1",
			Name:        "google pixel 3",
			Price:       22900.00,
			Amount:      3,
			DateCreated: now,
			DateUpdated: now,
		}
		newProduct := product.NewProduct{
			Name:   "google pixel 3",
			Price:  22900.00,
			Amount: 3,
		}

		actual, err := product.CreateNewProduct(db, newProduct, now)

		assert.Equal(t, nil, err)
		assert.Equal(t, expeted, actual)
	})
	t.Run("List All Product", func(t *testing.T) {
		now := time.Date(2019, time.June, 1, 0, 0, 0, 0, time.UTC)
		expeted := []product.Product{
			{
				ID:          "1",
				Name:        "google pixel 3",
				Price:       22900.00,
				Amount:      3,
				DateCreated: now,
				DateUpdated: now,
			},
		}

		actual, err := product.ListProduct(db)

		assert.Equal(t, nil, err)
		assert.Equal(t, expeted, actual)
	})
	t.Run("Get Product By ID", func(t *testing.T) {
		now := time.Date(2019, time.June, 1, 0, 0, 0, 0, time.UTC)
		expeted := product.Product{
			ID:          "1",
			Name:        "google pixel 3",
			Price:       22900.00,
			Amount:      3,
			DateCreated: now,
			DateUpdated: now,
		}

		actual, err := product.GetProductByID(db, "1")
		assert.Equal(t, nil, err)
		assert.Equal(t, expeted, actual)
	})
	t.Run("Update Product", func(t *testing.T) {
		now := time.Date(2019, time.June, 1, 0, 0, 0, 0, time.UTC)
		prod := product.UpdateProduct{
			Price: Float64Pointer(2300.00),
		}

		err := product.Update(db, "1", prod, now)
		assert.Equal(t, nil, err)
	})
	resual, err = database.DropTable(db, "products")
	if err != nil {
		t.Log("ceate tabel error: ", err)

	}
	t.Log(resual)
}

func Float64Pointer(f float64) *float64 {
	return &f
}
