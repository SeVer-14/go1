package service

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	dto "go1/internal/DTO"
	"go1/internal/entity"
	"go1/internal/repository"
)

type OrderService struct {
	repo        *repository.OrderRepository
	cartRepo    *repository.CartRepository
	productRepo *repository.ProductRepository
}

func NewOrderService(repo *repository.OrderRepository, cartRepo *repository.CartRepository, productRepo *repository.ProductRepository) *OrderService {
	return &OrderService{
		repo:        repo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *OrderService) buildOrderEntity(cartID uint, cartItems []entity.CartItem) (*entity.Order, float64, error) {
	order := &entity.Order{
		CartID:     cartID,
		Status:     "pending",
		OrderItems: make([]entity.OrderItem, 0, len(cartItems)),
	}

	var total float64
	for _, item := range cartItems {
		product, err := s.productRepo.GetByID(item.ProductID)
		if err != nil {
			return nil, 0, fmt.Errorf("product not found: %v", err)
		}

		order.OrderItems = append(order.OrderItems, entity.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
		total += float64(item.Quantity) * product.Price
	}
	order.Total = total

	return order, total, nil
}

func (s *OrderService) convertToCartDTO(cartID uint, items []entity.CartItem) (*dto.CartDTO, error) {
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

func (s *OrderService) convertToOrderDTO(order *entity.Order, total float64) *dto.OrderDTO {
	items := make([]dto.OrderItemDTO, len(order.OrderItems))
	for i, item := range order.OrderItems {
		items[i] = dto.OrderItemDTO{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}

	return &dto.OrderDTO{
		ID:     order.ID,
		Status: order.Status,
		Total:  total,
		Items:  items,
	}
}

func (s *OrderService) CreateOrder(cartID uint) (*dto.OrderDTO, error) {
	cartItems, err := s.cartRepo.GetCart(cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart: %v", err)
	}

	if len(cartItems) == 0 {
		return nil, errors.New("cart is empty")
	}

	orderEntity, total, err := s.buildOrderEntity(cartID, cartItems)
	if err != nil {
		return nil, fmt.Errorf("failed to build order: %v", err)
	}

	if err := s.repo.CreateOrder(orderEntity); err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	if err := s.cartRepo.ClearCart(cartID); err != nil {
		return nil, fmt.Errorf("failed to clear cart: %v", err)
	}

	return s.convertToOrderDTO(orderEntity, total), nil
}

func (s *OrderService) GetOrders(cartID uint) ([]dto.OrderDTO, error) {
	orders, err := s.repo.GetOrders(cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %v", err)
	}

	result := make([]dto.OrderDTO, len(orders))
	for i, order := range orders {
		result[i] = *s.convertToOrderDTO(&order, order.Total)
	}

	return result, nil
}

func (s *OrderService) UpdateOrderStatus(req dto.UpdateOrderStatusDTO) error {
	var validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	validStatuses := map[string]bool{
		"pending":    true,
		"processing": true,
		"completed":  true,
		"cancelled":  true,
	}

	if !validStatuses[req.Status] {
		return errors.New("invalid status")
	}

	return s.repo.UpdateOrderStatus(req.OrderID, req.Status)
}
