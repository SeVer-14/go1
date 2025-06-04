package service

import (
	"go1/database"
	"go1/entity"
	"gorm.io/gorm"
)

type ProductService interface {
	Add(product entity.Product) entity.Product
	Show() []entity.Product
	Delete(id uint) bool
	AddToCart(productID uint, userID uint) error
	GetCart(userID uint) ([]entity.Cart, error)
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
	err := database.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error
	return cartItems, err
}
