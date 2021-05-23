package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ksupdev/updev-go-ex-stock-api/db"
)

func Setup(router *gin.Engine) {

	db.SetupDB()

	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTransactionAPI(router)

}
