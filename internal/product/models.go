package product

// Product is a item is stroe
type Product struct {
	ID     string  `db:"product_id" json:"id"`
	Name   string  `db:"name" json:"name"`
	Price  float64 `db:"price" json:"price"`
	Amount int     `db:"amount" json:"amount"`
}
