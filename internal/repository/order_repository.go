package repository

import (
	"errors"
	"fmt"
	"go1/internal/entity"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(cart *entity.Cart) (*entity.Order, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order := &entity.Order{
		UserID: cart.UserID,
		Status: "pending",
		CartID: cart.ID,
	}

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	for _, item := range cart.CartItems {
		var product entity.Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("product not found: %v", err)
		}

		itemTotal := float64(item.Quantity) * product.Price
		order.Total += itemTotal

		orderItem := entity.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to create order item: %v", err)
		}
	}

	if err := tx.Model(order).Update("total", order.Total).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update order total: %v", err)
	}

	if err := tx.Where("cart_id = ?", cart.ID).Delete(&entity.CartItem{}).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to clear cart: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("transaction commit failed: %v", err)
	}

	var result entity.Order
	if err := r.db.
		Preload("OrderItems.Product").
		First(&result, order.ID).
		Error; err != nil {
		return nil, fmt.Errorf("failed to load order details: %v", err)
	}

	return &result, nil
}

func (r *OrderRepository) GetOrders(userID uint) ([]entity.Order, error) {
	var orders []entity.Order

	err := r.db.
		Preload("OrderItems.Product").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %v", err)
	}

	return orders, nil
}

func (r *OrderRepository) UpdateOrderStatus(orderID uint, status string) error {
	var order entity.Order
	if err := r.db.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("order not found")
		}
		return err
	}

	result := r.db.Model(&entity.Order{}).
		Where("id = ?", orderID).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected - order may not exist")
	}

	return nil
}
