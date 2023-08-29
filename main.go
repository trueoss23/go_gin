package main

import (
	"net/http"

	"go_gin/config"
	"go_gin/models"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/products", getProducts)

	router.GET("/product/:code", getProduct)

	router.POST("/products", addProduct)

	router.Run(config.Cfg.AppHost + ":" + config.Cfg.AppPort)
}

/*
	curl http://localhost:8083/products \
	    --include \
	    --header "Content-Type: application/json" \
	    --request "GET"
*/
func getProducts(c *gin.Context) {

	products := models.GetProducts()

	if products == nil || len(products) == 0 {

		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, products)

	}
}

/*
	curl http://localhost:8083/product/P0111 \
	    --include \
	    --header "Content-Type: application/json" \
	    --request "GET"
*/
func getProduct(c *gin.Context) {

	code := c.Param("code")

	product := models.GetProduct(code)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)

	} else {

		c.IndentedJSON(http.StatusOK, product)

	}

}

/*
	curl http://localhost:8083/products \
	    --include \
	    --header "Content-Type: application/json" \
	    --request "POST" \
	    --data '{"code": "P1114","name": "MacBook Air M1","qty": 10}'
*/
func addProduct(c *gin.Context) {

	var prod models.Product

	if err := c.BindJSON(&prod); err != nil {

		c.AbortWithStatus(http.StatusBadRequest)
	} else {

		models.AddProduct(prod)
		c.IndentedJSON(http.StatusCreated, prod)
	}

}
