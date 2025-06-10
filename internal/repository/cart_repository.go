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

func (r *CartRepository) AddItem(cartID uint, productID uint, quantity int) (*entity.CartItem, error) {
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
	if err := tx.Where("cart_Id = ?", cartID).FirstOrCreate(&cart, entity.Cart{CartID: cartID}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var item entity.CartItem
	if err := tx.Where("cart_Id = ? AND product_Id = ?", cart.CartID, productID).First(&item).Error; err == nil {
		item.Quantity += quantity
		if err := tx.Save(&item).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	} else {
		item = entity.CartItem{
			CartID:    cart.CartID,
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

func (r *CartRepository) GetCart(cartID uint) ([]entity.CartItem, error) {
	var items []entity.CartItem
	err := r.db.
		Where("cart_Id = ?", cartID).
		Find(&items).Error
	return items, err
}

func (r *CartRepository) RemoveItem(cartID uint, productID uint) error {
	var cart entity.Cart
	if err := r.db.Where("cart_Id = ?", cartID).First(&cart).Error; err != nil {
		return err
	}
	return r.db.Where("cart_Id = ? AND product_Id = ?", cart.ID, productID).Delete(&entity.CartItem{}).Error
}

func (r *CartRepository) ClearCart(cartID uint) error {
	if err := r.db.Where("cart_Id = ?", cartID).Delete(&entity.CartItem{}).Error; err != nil {
		return err
	}
	return nil
}
