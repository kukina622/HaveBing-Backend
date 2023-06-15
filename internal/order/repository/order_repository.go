package repository

import (
	"HaveBing-Backend/internal/domain"
	"context"
	"fmt"

	"gorm.io/gorm"
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
	var order domain.Order
	err := r.db.Joins("Payment").Joins("Shipping").Preload("User").Preload("OrderItem").Preload("OrderItem.Product.ProductCategory").First(&order, id).Error
	return &order, err
}

func (r *orderRepository) GetByUserId(ctx context.Context, userId uint) ([]domain.Order, error) {
	var orderList []domain.Order
	err := r.db.Joins("Payment").Joins("Shipping").Preload("User").Preload("OrderItem").Preload("OrderItem.Product.ProductCategory").Where("user_id = ?", userId).Find(&orderList).Error
	return orderList, err
}

func (r *orderRepository) Create(ctx context.Context, order *domain.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) Update(ctx context.Context, order *domain.Order) error {
	return nil
}
func (r *orderRepository) CreateWithTx(ctx context.Context, tx *gorm.DB, order *domain.Order) error {
	return tx.Create(order).Error
}

func (r *orderRepository) WithTransaction(ctx context.Context, txFunc func(*gorm.DB) error) (err error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = txFunc(tx)
	return
}
