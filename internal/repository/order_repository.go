package repository

import (
	"go1/internal/entity"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetOrdersWithItems(cartID uint) ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.
		Preload("OrderItems").
		Where("cart_Id = ?", cartID).
		Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) CreateOrder(order *entity.Order) error {
	return r.db.Create(order).Error
}
func (r *OrderRepository) GetOrders(cartID uint) ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.
		Preload("OrderItems").
		Where("cart_Id = ?", cartID).
		Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) UpdateOrderStatus(orderID uint, status string) error {
	return r.db.Model(&entity.Order{}).
		Where("id = ?", orderID).
		Update("status", status).Error
}
