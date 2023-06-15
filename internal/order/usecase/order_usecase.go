package usecase

import (
	"HaveBing-Backend/internal/domain"
	"HaveBing-Backend/internal/dto"
	"HaveBing-Backend/internal/util/snowflake"
	"context"
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderUseCase struct {
	orderRepo    domain.OrderRepository
	paymentRepo  domain.PaymentRepository
	shippingRepo domain.ShippingRepository
	productRepo  domain.ProductRepository
}

func New(
	orderRepo domain.OrderRepository,
	paymentRepo domain.PaymentRepository,
	shippingRepo domain.ShippingRepository,
	productRepo domain.ProductRepository,
) domain.OrderUseCase {
	return &OrderUseCase{
		orderRepo:    orderRepo,
		paymentRepo:  paymentRepo,
		shippingRepo: shippingRepo,
		productRepo:  productRepo,
	}
}

func (u *OrderUseCase) GetAll(ctx context.Context) ([]domain.Order, error) {
	return u.orderRepo.GetAll(ctx)
}

func (u *OrderUseCase) GetById(ctx context.Context, id uint) (*domain.Order, error) {
	return u.orderRepo.GetById(ctx, id)
}

func (u *OrderUseCase) GetByUserId(ctx context.Context, userId uint) ([]domain.Order, error) {
	return u.orderRepo.GetByUserId(ctx, userId)
}

func (u *OrderUseCase) Create(ctx context.Context, newOrder *dto.AddOrderDTO) (*domain.Order, error) {
	order := &domain.Order{
		OrderNumber: snowflake.GenerateID().String(),
		Status:      "preparing",
		UserId:      newOrder.UserId,
		Note:        newOrder.Note,
		OrderDate:   time.Now(),
		Email:       newOrder.Email,
	}

	err := u.orderRepo.WithTransaction(ctx, func(tx *gorm.DB) error {

		if newOrder.PaymentMethod != "transfer" {
			return errors.New("invalid payment method")
		}

		payment := domain.Payment{
			PaymentDate:   nil,
			ShippingFee:   decimal.NewFromInt(60),
			PaymentStatus: domain.PaymentStatusUnpaid,
			InvoiceType:   newOrder.InvoiceType,
			PaymentMethod: domain.PaymentMethodTransfer,
		}

		expectedDeliveryDate, err := time.Parse("2006-01-02", newOrder.ExpectedDeliveryDate)
		expectedDeliveryDateAddr := &expectedDeliveryDate

		if err != nil {
			expectedDeliveryDateAddr = nil
		}

		shipping := domain.Shipping{
			RecipientName:        newOrder.RecipientName,
			RecipientPhone:       newOrder.RecipientPhone,
			Address:              newOrder.Address,
			ShippingMethod:       newOrder.ShippingMethod,
			ExpectedDeliveryDate: expectedDeliveryDateAddr,
		}

		orderItemList := []domain.OrderItem{}
		amount := decimal.NewFromInt(0)

		for _, item := range newOrder.ProductList {
			var orderItem domain.OrderItem
			if err := u.productRepo.DecreaseInventoryWithTx(ctx, tx, item.ProductId, item.Quantity); err != nil {
				return err
			}

			product, err := u.productRepo.GetById(ctx, item.ProductId)
			if err != nil {
				return err
			}

			orderItem.Product = *product
			orderItem.Quality = item.Quantity
			orderItem.Price = product.Price
			orderItemList = append(orderItemList, orderItem)

			amount = amount.Add(product.Price.Mul(decimal.NewFromInt(int64(item.Quantity))))
		}

		payment.Amount = amount
		order.OrderItem = orderItemList
		order.Payment = payment
		order.Shipping = shipping

		if err := u.orderRepo.CreateWithTx(ctx, tx, order); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	currentOrder, err := u.orderRepo.GetById(ctx, order.ID)
	if err != nil {
		return nil, err
	}
	return currentOrder, nil
}

func (u *OrderUseCase) Update(ctx context.Context, order *domain.Order) error {
	return nil
}
