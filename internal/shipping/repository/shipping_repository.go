package repository

import (
	"HaveBing-Backend/internal/domain"
	"context"

	"gorm.io/gorm"
)

type ShippingRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ShippingRepository {
	return &ShippingRepository{db: db}
}

func (r *ShippingRepository) GetAll(ctx context.Context) ([]domain.Shipping, error) {
	return nil, nil
}

func (r *ShippingRepository) GetById(ctx context.Context, id uint) (*domain.Shipping, error) {
	return nil, nil
}

func (r *ShippingRepository) Create(ctx context.Context, shipping *domain.Shipping) error {
	return nil
}

func (r *ShippingRepository) Update(ctx context.Context, shipping *domain.Shipping) error {
	return nil
}
