package controller

import (
	"go1/internal/service"
)

type ProductController interface {
}

type controller struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &controller{service: service}
}
