package api

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ksupdev/updev-go-ex-stock-api/db"
	"github.com/ksupdev/updev-go-ex-stock-api/interceptor"
	"github.com/ksupdev/updev-go-ex-stock-api/model"
)

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		//productAPI.GET("/product", myInterceptor, myInterceptor, myInterceptor, getProduct)
		productAPI.GET("/product", interceptor.JwtVerify, getProduct)
		productAPI.GET("/product/:id", getProductByID)
		productAPI.POST("/product", createProduct)
		productAPI.PUT("/product", editProduct)
	}
}

func getProduct(c *gin.Context) {
	var product []model.Product
	if keyword := c.Query("keyword"); isEmptyString(keyword) {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("name like ?", keyword).Find(&product)
	} else {
		db.GetDB().Find(&product)
	}
	c.JSON(http.StatusOK, gin.H{"action": "get product", "data": product})
}

func getProductByID(c *gin.Context) {
	var product []model.Product
	if id, ok := c.Params.Get("id"); ok {
		db.GetDB().Where("id = ?", id).First(&product)
	}
	c.JSON(http.StatusOK, gin.H{"action": "get product by id", "data": product})
}

func createProduct(c *gin.Context) {

	var product model.Product
	// This method need to implement validate requir field
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64) // it will return 0 , When convertion value doesn't success
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	product.CreateAt = time.Now()
	db.GetDB().Create(&product)
	image, _ := c.FormFile("image")
	saveImage(image, &product, c)
	c.JSON(http.StatusOK, gin.H{"result": product})
}

func editProduct(c *gin.Context) {
	var product model.Product
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)
	product.ID = uint(id)
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

	db.GetDB().Save(&product)
	image, _ := c.FormFile("image")
	saveImage(image, &product, c)
	c.JSON(http.StatusOK, gin.H{"result": product})
}

func saveImage(image *multipart.FileHeader, product *model.Product, c *gin.Context) {
	if image != nil {
		runningDir, _ := os.Getwd()
		product.Image = image.Filename
		extension := filepath.Ext(image.Filename)
		fileName := fmt.Sprintf("%d%s", product.ID, extension)
		filePath := fmt.Sprintf("%s/uploaded/images/%s", runningDir, fileName)

		if fileExists(filePath) {
			os.Remove(filePath)
		}
		c.SaveUploadedFile(image, filePath)
		db.GetDB().Model(&product).Update("image", fileName)

	}

}

func fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isEmptyString(data string) bool {
	return len(strings.TrimSpace(data)) <= 0
}
