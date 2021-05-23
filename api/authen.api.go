package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksupdev/updev-go-ex-stock-api/model"
)

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}

}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "login"})
}

func register(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		c.JSON(http.StatusOK, gin.H{"result": "register", "data": user})
	}

}
