package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CartID    uint       `gorm:"unique"`
	CartItems []CartItem `gorm:"foreignKey:CartID;references:CartID"`
}

type CartItem struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductID"`
}
