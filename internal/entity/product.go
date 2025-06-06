package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductID int     `json:"ProductID" gorm:"unique"`
	Title     string  `json:"title" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}
