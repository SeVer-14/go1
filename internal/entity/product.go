package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductID int `gorm:"unique"`
	Title     string
	Price     float64
}
