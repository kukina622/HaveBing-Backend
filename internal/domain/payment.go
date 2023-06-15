package domain

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type paymentStatus string
type PaymentMethod string

const (
	PaymentStatusUnpaid   paymentStatus = "unpaid"
	PaymentStatusPaid     paymentStatus = "paid"
	PaymentMethodTransfer PaymentMethod = "transfer"
)

type Payment struct {
	ID            uint            `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"paymentId"`
	PaymentDate   *time.Time      `gorm:"type:datetime" json:"paymentDate"`
	ShippingFee   decimal.Decimal `gorm:"type:decimal(19,4) NOT NULL;" json:"shippingFee"`
	Amount        decimal.Decimal `gorm:"type:decimal(19,4) NOT NULL;" json:"amount"`
	PaymentMethod PaymentMethod   `gorm:"type:ENUM('transfer') NOT NULL;" json:"paymentMethod"`
	InvoiceType   string          `gorm:"type:varchar(255) NOT NULL;" json:"invoiceType"`
	PaymentStatus paymentStatus   `gorm:"type:ENUM('unpaid', 'paid') NOT NULL;" json:"paymentStatus"`
	OrderId       uint            `gorm:"type:bigint(20) NOT NULL;index:idx_order_id;" json:"orderId"`
}

type PaymentRepository interface {
	GetById(ctx context.Context, id uint) (*Payment, error)
	Create(ctx context.Context, payment *Payment) error
	Update(ctx context.Context, payment *Payment) error
}
