package repository

import (
	"HaveBing-Backend/internal/domain"
	"context"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) GetAll(ctx context.Context) ([]domain.Payment, error) {
	return nil, nil
}

func (r *PaymentRepository) GetById(ctx context.Context, id uint) (*domain.Payment, error) {
	return nil, nil
}

func (r *PaymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	return nil
}

func (r *PaymentRepository) Update(ctx context.Context, payment *domain.Payment) error {
	return nil
}
