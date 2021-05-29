package interceptor

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ksupdev/updev-go-ex-stock-api/model"
)

var secreatKey = "87654321"

func JwtSign(payload model.User) string {
	atClaims := jwt.MapClaims{}

	// Payload begin
	atClaims["id"] = payload.ID
	atClaims["username"] = payload.Username
	atClaims["level"] = payload.Level
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// Payload end

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secreatKey))

	return token

}

type MyError struct{}

func (m *MyError) Error() string {
	return "boom"
}

func JwtVerify(c *gin.Context) {

	if headerAuthoriza := c.GetHeader("Authorization"); headerAuthoriza == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"result": "nok", "message": " Authorization code not found"})
		c.Abort()
	} else {
		// tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
		tokenString := strings.Split(headerAuthoriza, " ")[1]
		fmt.Println(tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				//return nil, fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
				//return nil, &MyError{}
				//fmt.Printf("Show payload : %v \n", token.Claims)
				errorMessage := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
				return nil, errors.New(errorMessage)
			}
			return []byte(secreatKey), nil
		})

		// Cast token,Claims with jwt.MapClaims => token.Claims.(jwt.MapClaims)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			staffID := fmt.Sprintf("%v", claims["id"])
			username := fmt.Sprintf("%v", claims["username"])
			level := fmt.Sprintf("%v", claims["level"])

			// Get jwt data then set thems to Request context
			c.Set("jwt_staff_id", staffID)
			c.Set("jwt_username", username)
			c.Set("jwt_level", level)

			// You can get these value with
			/*
				c.GetString("jwt_staff_id")
				c.GetString("jwt_username")
				c.GetString("jwt_level")

				the c is Context
			*/

			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "nok", "message": "invalid token", "error": err})
			c.Abort()
		}
	}

}
