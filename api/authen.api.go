package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ksupdev/updev-go-ex-stock-api/db"
	"github.com/ksupdev/updev-go-ex-stock-api/interceptor"
	"github.com/ksupdev/updev-go-ex-stock-api/model"
	"golang.org/x/crypto/bcrypt"
)

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}

}

func login(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		var queryUser model.User
		if err := db.GetDB().First(&queryUser, "username = ?", user.Username).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "nok", "error": err})
		} else if !checkPasswordHash(user.Password, queryUser.Password) {
			c.JSON(http.StatusOK, gin.H{"result": "nok", "error": "invalid password"})
		} else {
			token := interceptor.JwtSign(queryUser)
			c.JSON(http.StatusOK, gin.H{"result": "OK", "token": token})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Unable to bind data"})
	}

}

func register(c *gin.Context) {
	var user model.User
	var hashError error
	if c.ShouldBind(&user) == nil {

		user.Password, hashError = hashPassword(user.Password)
		if hashError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Unable to hash password"})
		}

		user.CreateAt = time.Now()

		if err := db.GetDB().Create(&user).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "nok", "error": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "ok", "data": user})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Unable to bind data"})
	}
}

func hashPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(byte), err
}

func checkPasswordHash(password, hash string) bool {
	error := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return error == nil
}
