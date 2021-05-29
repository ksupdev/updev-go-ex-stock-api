package db

import (
	"github.com/ksupdev/updev-go-ex-stock-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func SetupDB() {
	/*
		==> Connect SQLlight ("gorm.io/driver/sqlite")
			database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	*/

	// Connect to mysql
	dsn := "root:rootpw@tcp()/cmgosstock?parseTime=true&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Faild to connect database")
	}

	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.Transaction{})

	db = database

}
