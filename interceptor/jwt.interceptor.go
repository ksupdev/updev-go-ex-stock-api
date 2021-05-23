package interceptor

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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

func JwtVerify() {}
