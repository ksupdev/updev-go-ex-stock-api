# updev-go-ex-stock-api
go workshop

## setup module
go mod init github.com/ksupdev/updev-go-ex-stock-api
go get github.com/gin-gonic/gin

### driver for SQLite
go get gorm.io/driver/sqlite

### driver for MYSQL
go get gorm.io/driver/mysql

### driver for GORM
go get gorm.io/gorm

go get github.com/dgrijalva/jwt-go

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

- Auto create table structure incase you don't have that table before
```golang
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.Transaction{})
```
> Gorm auto create table name User => users


> strconv.ParseInt(c.PostForm("stock"), 10, 64) กรณีที่ไม่สามารถ convert ได้จะ return 0 มาให้เลย

## Build to production
- Simple build ``go build`` ซึ่งจะได้ file ที่มีชื่อเดียวกับ module โดยการ build แบบนี้จะมีขนาดใหญ่มาก (14.7 mb)

```powershell
	go build
	./updev-go-ex-stock-api
```

- Build and remove the junk file ``go build -ldflags "-s -w"`` จะได้ file ที่มีขนาดเล็กกว่าแบบแรก (11.6 mb)
```powershell
	go build -ldflags "-s -w"
	./updev-go-ex-stock-api
```




