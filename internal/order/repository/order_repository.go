package repository

import (
	"HaveBing-Backend/internal/domain"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetAll(ctx context.Context) ([]domain.Order, error) {
	var orderList []domain.Order
	err := r.db.Find(&orderList).Error
	return orderList, err
}

func (r *orderRepository) GetById(ctx context.Context, id uint) (*domain.Order, error) {
	return nil, nil
}

func (r *orderRepository) GetByUserId(ctx context.Context, userId uint) ([]domain.Order, error) {
	var orderList []domain.Order
	err := r.db.Preload(clause.Associations).Where("user_id = ?", userId).Find(&orderList).Error
	return orderList, err
}

func (r *orderRepository) Create(ctx context.Context, order *domain.Order) error {
	return nil
}

func (r *orderRepository) Update(ctx context.Context, order *domain.Order) error {
	return nil
}
