package handlers

import (
	"log"
	"net/http"
	"smalldoc124/service/internal/product"

	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	DB        *sqlx.DB
	ProductDB product.ProductDB
}

func (prod Product) GetProductByID(context *gin.Context) {
	productID := context.Param("id")

	product, err := prod.ProductDB.GetProductByID(productID, prod.DB)
	if err != nil {
		log.Println("Handlers GetProductByID error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, product)
}
