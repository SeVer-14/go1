package service

import (
	"go1/internal/entity"
	"go1/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}
func (s *OrderService) CreateOrder(cart *entity.Cart) (*entity.Order, error) {
	return s.repo.CreateOrder(cart)
}

func (s *OrderService) GetOrders(userID uint) ([]entity.Order, error) {
	return s.repo.GetOrders(userID)
}

func (s *OrderService) UpdateOrderStatus(orderID uint, status string) error {
	return s.repo.UpdateOrderStatus(orderID, status)
}
