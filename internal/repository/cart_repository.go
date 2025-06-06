package repository

import (
	"errors"
	"go1/internal/entity"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) AddItem(userID uint, productID uint, quantity int) (*entity.CartItem, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var product entity.Product
	if err := tx.First(&product, productID).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("product not found")
	}

	var cart entity.Cart
	if err := tx.Where("user_id = ?", userID).FirstOrCreate(&cart, entity.Cart{UserID: userID}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var item entity.CartItem
	if err := tx.Where("cart_id = ? AND product_id = ?", cart.ID, productID).First(&item).Error; err == nil {
		item.Quantity += quantity
		if err := tx.Save(&item).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	} else {
		item = entity.CartItem{
			CartID:    cart.ID,
			ProductID: productID,
			Quantity:  quantity,
		}
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *CartRepository) GetCart(userID uint) (*entity.Cart, error) {
	var cart entity.Cart
	err := r.db.
		Preload("CartItems.Product"). // Загружаем связанные продукты
		Where("user_id = ?", userID).
		First(&cart).Error
	return &cart, err
}

func (r *CartRepository) RemoveItem(userID uint, productID uint) error {
	var cart entity.Cart
	if err := r.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return err
	}
	return r.db.Where("cart_id = ? AND product_id = ?", cart.ID, productID).Delete(&entity.CartItem{}).Error
}
