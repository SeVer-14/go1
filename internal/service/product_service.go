package service

import (
	"errors"
	dto "go1/internal/DTO"
	"go1/internal/entity"
	"go1/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]dto.ProductDTO, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var response []dto.ProductDTO
	for _, p := range products {
		response = append(response, dto.ProductDTO{
			ID:        p.ID,
			ProductID: p.ProductID,
			Title:     p.Title,
			Price:     p.Price,
		})
	}
	return response, nil
}

func (s *ProductService) CreateProduct(req dto.ProductDTO) (*dto.ProductDTO, error) {
	if req.Price <= 0 {
		return nil, errors.New("price must be positive")
	}

	product := &entity.Product{
		ProductID: req.ProductID,
		Title:     req.Title,
		Price:     req.Price,
	}

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}

	return &dto.ProductDTO{
		ID:        product.ID,
		ProductID: product.ProductID,
		Title:     product.Title,
		Price:     product.Price,
	}, nil
}

func (s *ProductService) UpdateProduct(id uint, input dto.ProductDTO) (*dto.ProductDTO, error) {
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
