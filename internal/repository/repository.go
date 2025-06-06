package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	Product *ProductRepository
	Cart    *CartRepository
	Order   *OrderRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Product: NewProductRepository(db),
		Cart:    NewCartRepository(db),
		Order:   NewOrderRepository(db),
	}
}
