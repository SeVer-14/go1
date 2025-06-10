package service

import (
	"errors"
	"fmt"
	dto "go1/internal/DTO"
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

func (s *CartService) AddToCart(cartID uint, req dto.AddToCartDTO) (*dto.CartItemDTO, error) {
	if req.Quantity <= 0 {
		return nil, errors.New("quantity must be positive")
	}

	product, err := s.productRepo.GetByID(req.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	item, err := s.repo.AddItem(cartID, req.ProductID, req.Quantity)
	if err != nil {
		return nil, err
	}

	return &dto.CartItemDTO{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		Price:     product.Price,
		Title:     product.Title,
	}, nil
}

func (s *CartService) GetCart(cartID uint) (*dto.CartDTO, error) {
	items, err := s.repo.GetCart(cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart items: %v", err)
	}

	cartDTO := &dto.CartDTO{
		CartID: cartID,
		Items:  make([]dto.CartItemDTO, 0, len(items)),
	}

	var total float64
	for _, item := range items {
		product, err := s.productRepo.GetByID(item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product %d: %v", item.ProductID, err)
		}

		cartDTO.Items = append(cartDTO.Items, dto.CartItemDTO{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
			Title:     product.Title,
		})

		total += float64(item.Quantity) * product.Price
	}
	cartDTO.Total = total

	return cartDTO, nil
}

func (s *CartService) RemoveFromCart(cartID uint, productID uint) error {
	return s.repo.RemoveItem(cartID, productID)
}
