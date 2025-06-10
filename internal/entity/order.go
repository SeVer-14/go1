package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CartID     uint        `gorm:"index"`
	Status     string      `gorm:"default:'pending'"` //pending processing completed cancelled
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	Total      float64     `gorm:"type:decimal(10,2)"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint `gorm:"index"`
	ProductID uint `gorm:"index"`
	Quantity  int
	Price     float64 `gorm:"type:decimal(10,2)"`
}
