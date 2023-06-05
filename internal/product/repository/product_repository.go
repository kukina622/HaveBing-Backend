package repository

import (
	"HaveBing-Backend/internal/domain"
	"context"

	"gorm.io/gorm"
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
	return nil, nil
}

func (p *ProductRepository) GetById(ctx context.Context, id uint) (*domain.Product, error) {
	return nil, nil
}

func (p *ProductRepository) GetByCategoryId(ctx context.Context, categoryId uint) ([]domain.Product, error) {
	return nil, nil
}

func (p *ProductRepository) Save(ctx context.Context, product *domain.Product) error {
	return nil
}
