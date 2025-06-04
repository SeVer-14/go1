package service

import (
	"errors"
	"go1/database"
	"go1/entity"
	"gorm.io/gorm"
	"time"
)

type ProductService interface {
	Add(product entity.Product) entity.Product
	Show() []entity.Product
	Delete(id uint) bool
	AddToCart(productID uint, userID uint) error
	GetCart(userID uint) ([]entity.Cart, error)
	RemoveFromCart(cartID uint, userID uint) error
	UpdateCartItem(cartID uint, userID uint, quantity int) error
	CreateOrder(userID uint) (entity.Order, error)
	GetOrders(userID uint) ([]entity.Order, error)
	UpdateOrderStatus(orderID uint, status string) error
}

type productService struct {
}

func NewProductService() ProductService {
	return &productService{}
}

func (s *productService) Add(product entity.Product) entity.Product {
	database.DB.Create(&product)
	return product
}

func (s *productService) Show() []entity.Product {
	var products []entity.Product
	database.DB.Find(&products)
	return products
}

func (s *productService) Delete(id uint) bool {
	result := database.DB.Delete(&entity.Product{}, id)
	return result.RowsAffected > 0
}

func (s *productService) AddToCart(productID uint, userID uint) error {
	var cart entity.Cart
	if err := database.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			cart = entity.Cart{
				UserID:    userID,
				ProductID: productID,
				Quantity:  1,
			}
			return database.DB.Create(&cart).Error
		}
		return err
	}
	cart.Quantity++
	return database.DB.Save(&cart).Error
}

func (s *productService) GetCart(userID uint) ([]entity.Cart, error) {
	var cartItems []entity.Cart
	err := database.DB.
		Preload("Product").
		Where("user_id = ?", userID).
		Find(&cartItems).Error
	return cartItems, err
}

func (s *productService) RemoveFromCart(cartID uint, userID uint) error {
	result := database.DB.Where("id = ? AND user_id = ?", cartID, userID).Delete(&entity.Cart{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("cart item not found")
	}
	return nil
}

func (s *productService) UpdateCartItem(cartID uint, userID uint, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	result := database.DB.Model(&entity.Cart{}).
		Where("id = ? AND user_id = ?", cartID, userID).
		Update("quantity", quantity)

	return result.Error
}
func (s *productService) CreateOrder(userID uint) (entity.Order, error) {
	var order entity.Order
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var cartItems []entity.Cart
		if err := tx.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
			return err
		}

		if len(cartItems) == 0 {
			return errors.New("cart is empty")
		}

		var total float64
		for _, item := range cartItems {
			total += item.Product.Price * float64(item.Quantity)
		}

		order = entity.Order{
			UserID:    userID,
			Total:     total,
			Status:    "created",
			CreatedAt: time.Now(),
		}
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		var orderItems []entity.OrderItem
		for _, item := range cartItems {
			orderItem := entity.OrderItem{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Product.Price,
			}
			if err := tx.Create(&orderItem).Error; err != nil {
				return err
			}
			orderItems = append(orderItems, orderItem)
		}

		if err := tx.Where("user_id = ?", userID).Delete(&entity.Cart{}).Error; err != nil {
			return err
		}

		if err := tx.Preload("Product").Where("order_id = ?", order.ID).Find(&orderItems).Error; err != nil {
			return err
		}
		order.Items = orderItems

		return nil
	})

	return order, err
}

func (s *productService) GetOrders(userID uint) ([]entity.Order, error) {
	var orders []entity.Order
	err := database.DB.
		Preload("Items").
		Preload("Items.Product").
		Where("user_id = ?", userID).
		Find(&orders).Error
	return orders, err
}

func (s *productService) UpdateOrderStatus(orderID uint, status string) error {
	validStatuses := map[string]bool{
		"created":    true,
		"processing": true,
		"completed":  true,
		"cancelled":  true,
	}

	if !validStatuses[status] {
		return errors.New("invalid status")
	}

	return database.DB.Model(&entity.Order{}).Where("id = ?", orderID).Update("status", status).Error
}
