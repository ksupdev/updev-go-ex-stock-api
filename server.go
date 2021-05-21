package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi this Stock system")
	router := gin.Default()
	router.Static("/images", "./uploaded/images")
	router.Run(":8081")

}
