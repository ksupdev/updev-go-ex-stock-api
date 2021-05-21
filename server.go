package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ksupdev/updev-go-ex-stock-api/api"
)

func main() {
	fmt.Println("Hi this Stock system")
	router := gin.Default()
	api.Setup(router)

	router.Static("/images", "./uploaded/images")
	router.Run(":8081")

}
