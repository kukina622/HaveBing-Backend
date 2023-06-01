package repository

import (
	"HaveBing-Backend/internal/domain"
	"context"
	"gorm.io/gorm"
)

type ProductCategoryRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProductCategoryRepository {
	return &ProductCategoryRepository{
		db: db,
	}
}

func (p *ProductCategoryRepository) GetAll(ctx context.Context) ([]domain.ProductCategory, error) {
	var productCategoryList []domain.ProductCategory
	result := p.db.Find(&productCategoryList)
	return productCategoryList, result.Error
}

func (p *ProductCategoryRepository) GetById(ctx context.Context, id int) (*domain.ProductCategory, error) {
	var productCategory domain.ProductCategory
	err := p.db.First(&productCategory, id).Error
	return &productCategory, err
}

func (p *ProductCategoryRepository) Update(ctx context.Context, productCategory *domain.ProductCategory) error {
	return nil
}

func (p *ProductCategoryRepository) Save(ctx context.Context, productCategory *domain.ProductCategory) error {
	return p.db.Create(productCategory).Error
}
