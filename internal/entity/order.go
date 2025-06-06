package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint        `json:"user_id" gorm:"index"`
	Status     string      `json:"status" gorm:"default:'pending'"` //required processing completed cancelled
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	Total      float64     `json:"total" gorm:"type:decimal(10,2)"`
	CartID     uint        `json:"-" gorm:"index"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id" gorm:"index"`
	ProductID uint    `json:"product_id" gorm:"index"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2)"`
}
