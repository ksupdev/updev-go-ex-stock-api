package model

import (
	"time"
)

type Product struct {
	ID uint `gorm:"primary_key"`
	//Name string `form:"name" binding:"required"`
	Name     string
	Stock    int64
	Price    float64
	Image    string
	CreateAt time.Time
}

func (p Product) IsValid() (errors []error) {
	var setOfErrors []error = nil
	//if len(strings.TrimSpace(p.Name)) <= 0 {
	//setOfErrors = append(setOfErrors, "A valid Product name is required %w")
	//}

	return setOfErrors
}
