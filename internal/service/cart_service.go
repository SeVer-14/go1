package service

import (
	"go1/internal/entity"
	"go1/internal/repository"
)

type CartService struct {
	repo        *repository.CartRepository
	productRepo *repository.ProductRepository
}

func NewCartService(cartRepo *repository.CartRepository, productRepo *repository.ProductRepository) *CartService {
	return &CartService{
		repo:        cartRepo,
		productRepo: productRepo,
	}
}

func (s *CartService) AddToCart(userID uint, productID uint, quantity int) (*entity.CartItem, error) {
	return s.repo.AddItem(userID, productID, quantity)
}

func (s *CartService) GetCart(userID uint) (*entity.Cart, error) {
	return s.repo.GetCart(userID)
}

func (s *CartService) RemoveFromCart(userID uint, productID uint) error {
	return s.repo.RemoveItem(userID, productID)
}
