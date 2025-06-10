package repository

import (
	dto "go1/internal/DTO"
	"go1/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	if err := r.db.Table("products").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) Create(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) Update(product *dto.ProductDTO) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Product{}, id).Error
}

func (r *ProductRepository) GetByID(id uint) (*dto.ProductDTO, error) {
	var product dto.ProductDTO
	if err := r.db.Table("products").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
