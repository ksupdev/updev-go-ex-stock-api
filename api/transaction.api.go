package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ksupdev/updev-go-ex-stock-api/db"
	"github.com/ksupdev/updev-go-ex-stock-api/interceptor"
	"github.com/ksupdev/updev-go-ex-stock-api/model"
)

func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("/api/v2")
	{
		transactionAPI.GET("/transaction", getTransaction)
		transactionAPI.POST("/transaction", interceptor.JwtVerify, createTransaction)
	}

}

func getTransaction(c *gin.Context) {
	var result []TransactionResult
	db.GetDB().Debug().Raw("SELECT payload, payload CustPayload, transactions.id, total, paid, payment_type, payment_detail, order_list, users.username Staff, transactions.created_at FROM transactions join users on transactions.staff_id = users.id", nil).Scan(&result)
	var temp model.Transaction
	json.Unmarshal([]byte(result[1].Payload), &temp)
	fmt.Println("data = ", temp.ID, " staff = ", temp.StaffID, " paid = ", temp.Paid)
	c.JSON(http.StatusOK, gin.H{"status": "Get transaction", "datas": result})
}

// This method need to add validation logic
func createTransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBind(&transaction); err == nil {
		transaction.StaffID = c.GetString("jwt_staff_id")
		transaction.CreatedAt = time.Now()
		pay, _ := json.Marshal(transaction)
		transaction.Payload = string(pay)
		db.GetDB().Create(&transaction)
		c.JSON(http.StatusOK, gin.H{"result": "ok", "data": transaction})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"result": "nok"})
	}

}

type TransactionResult struct {
	ID            uint
	Total         float64
	Paid          float64
	Change        float64
	PaymentType   string
	PaymentDetail string
	OrderList     string
	Staff         string
	Payload       string
	CustPayload   string
	CreateAt      time.Time
}
