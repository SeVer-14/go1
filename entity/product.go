package entity

import (
	"time"
)

type Product struct {
	ID        uint    `gorm:"primaryKey"`
	ProductID uint    `gorm:"unique"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
}

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID;references:ProductID"`
	Quantity  int     `json:"quantity"`
}

type Order struct {
	ID        uint `gorm:"primaryKey" json:"ID"`
	UserID    uint
	Items     []OrderItem `gorm:"foreignKey:OrderID;references:ID" json:"Items"`
	Total     float64
	Status    string
	CreatedAt time.Time
}

type OrderItem struct {
	ID        uint `gorm:"primaryKey" json:"ID"`
	OrderID   uint
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID;references:ProductID" json:"Product"`
	Quantity  int
	Price     float64
}
