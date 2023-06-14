package domain

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type Payment struct {
	ID            uint            `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"paymentId"`
	PaymentDate   *time.Time      `gorm:"type:datetime"`
	ShippingFee   decimal.Decimal `gorm:"type:decimal(19,4) NOT NULL;"`
	Amount        decimal.Decimal `gorm:"type:decimal(19,4) NOT NULL;"`
	PaymentStatus string          `gorm:"type:varchar(255) NOT NULL;"`
	OrderId       uint            `gorm:"type:bigint(20) NOT NULL;index:idx_order_id;" json:"orderId"`
}

type PaymentRepository interface {
	GetById(ctx context.Context, id uint) (*Payment, error)
	Create(ctx context.Context, payment *Payment) error
	Update(ctx context.Context, payment *Payment) error
}
