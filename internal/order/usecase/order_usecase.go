package usecase

import (
	"HaveBing-Backend/internal/domain"
	"context"
)

type OrderUseCase struct {
	orderRepo    domain.OrderRepository
	paymentRepo  domain.PaymentRepository
	shippingRepo domain.ShippingRepository
}

func New(orderRepo domain.OrderRepository, paymentRepo domain.PaymentRepository, shippingRepo domain.ShippingRepository) domain.OrderUseCase {
	return &OrderUseCase{
		orderRepo:    orderRepo,
		paymentRepo:  paymentRepo,
		shippingRepo: shippingRepo,
	}
}

func (u *OrderUseCase) GetAll(ctx context.Context) ([]domain.Order, error) {
	return u.orderRepo.GetAll(ctx)
}

func (u *OrderUseCase) GetById(ctx context.Context, id uint) (*domain.Order, error) {
	return nil, nil
}

func (u *OrderUseCase) GetByUserId(ctx context.Context, userId uint) ([]domain.Order, error) {
	return u.orderRepo.GetByUserId(ctx, userId)
}

func (u *OrderUseCase) Create(ctx context.Context, order *domain.Order) error {
	return nil
}

func (u *OrderUseCase) Update(ctx context.Context, order *domain.Order) error {
	return nil
}
