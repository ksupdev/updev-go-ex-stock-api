package model

import "time"

// https://gorm.io/docs/models.html

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique" form:"username" binding:"required"`
	// form => json username map to Username and this field is requitre
	Password string `gorm:"password" binding:"required"`
	Level    string `gorm:"default:normal"`
	CreateAt time.Time
}
