package service

import (
	"go1/internal/entity"
	"go1/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}
type productService struct {
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]entity.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) CreateProduct(input entity.Product) (*entity.Product, error) {
	product := &entity.Product{
		ProductID: input.ProductID,
		Title:     input.Title,
		Price:     input.Price,
	}

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) UpdateProduct(id uint, input entity.Product) (*entity.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	product.ProductID = input.ProductID
	product.Title = input.Title
	product.Price = input.Price

	if err := s.repo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
