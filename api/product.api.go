package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", getProduct)
		productAPI.GET("/product/:id", getProductByID)
		productAPI.POST("/product", createProduct)
		productAPI.PUT("/product", editProduct)
	}
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "get product"})
}

func getProductByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "get product by ID"})
}

func createProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "create product"})
}

func editProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "edit product"})
}
