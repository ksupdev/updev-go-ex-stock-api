# updev-go-ex-stock-api
go workshop

## setup module
go mod init github.com/ksupdev/updev-go-ex-stock-api
go get github.com/gin-gonic/gin

go get gorm.io/driver/sqlite
go get gorm.io/gorm

## Feature

- get images => http://localhost:8081/images/go.png


## noted
https://gorm.io/docs/models.html

- Method ``func (*gin.Context).ShouldBind(obj interface{}) error`` use for binding the JSON value from http request context to Struct. 

```golang
func register(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		c.JSON(http.StatusOK, gin.H{"result": "register", "data": user})
	}

}
```