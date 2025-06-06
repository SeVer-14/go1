package service

import (
	"go1/internal/repository"
)

type Services struct {
	Product *ProductService
	Cart    *CartService
	Order   *OrderService
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		Product: NewProductService(deps.Repos.Product),
		Cart:    NewCartService(deps.Repos.Cart, deps.Repos.Product),
		Order:   NewOrderService(deps.Repos.Order),
	}
}
