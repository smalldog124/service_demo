package tests

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http/httptest"
	"smalldoc124/service/cmd/sales-api/handlers"
	"smalldoc124/service/internal/product"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func Test_GetProductByID_Input_ID_1_Should_Mobile_Phone(t *testing.T) {
	expeted := `{"id":"1","name":"sony xperia","price":9999,"amount":5,"date_created":"0001-01-01T00:00:00Z","date_updated":"0001-01-01T00:00:00Z"}`

	request := httptest.NewRequest("GET", "/v1/product/1", nil)
	writer := httptest.NewRecorder()
	mockDBProductDB := new(mockDBProduct)
	mockDBProductDB.On("GetProductByID",mock.Anything,"1").Return(product.Product{ID: "1", Name: "sony xperia", Price: 9999.00, Amount: 5}, nil)
	product := handlers.Product{
		DB:&sqlx.DB{},
		ProductDB: mockDBProductDB,
	}

	testRoute := gin.Default()
	testRoute.GET("/v1/product/:id", product.GetProductByID)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actualtProduct, err := ioutil.ReadAll(response.Body)

	assert.Equal(t, nil, err)
	assert.Equal(t, expeted, string(actualtProduct))
}
