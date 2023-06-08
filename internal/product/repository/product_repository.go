package repository

import (
	"HaveBing-Backend/internal/domain"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var productList []domain.Product
	err := p.db.Preload(clause.Associations).Find(&productList).Error
	return productList, err
}

func (p *ProductRepository) GetById(ctx context.Context, id uint) (*domain.Product, error) {
	var product domain.Product
	err := p.db.Preload(clause.Associations).First(&product, id).Error
	return &product, err
}

func (p *ProductRepository) GetByCategoryId(ctx context.Context, categoryId uint) ([]domain.Product, error) {
	return nil, nil
}

func (p *ProductRepository) Save(ctx context.Context, product *domain.Product) error {
	return nil
}
