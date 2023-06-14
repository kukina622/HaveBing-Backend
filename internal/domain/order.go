package domain

import (
	"HaveBing-Backend/internal/dto"
	"context"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusShipping  OrderStatus = "shipping"
	OrderStatusDone      OrderStatus = "done"
	OrderStatusCanceled  OrderStatus = "canceled"
)

type Order struct {
	ID          uint        `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"orderId"`
	OrderNumber string      `gorm:"type:varchar(255) NOT NULL;"`
	Status      OrderStatus `gorm:"type:ENUM('preparing', 'shipping', 'done', 'canceled') NOT NULL;"`
	UserId      uint        `gorm:"type:bigint(20) NOT NULL;index:idx_user_id;" json:"userId"`
	Note        string      `gorm:"type:longtext"`
	OrderDate   time.Time   `gorm:"type:datetime"`
	User        User        `gorm:"foreignKey:UserId"`
	Payment     Payment     `gorm:"foreignKey:OrderId"`
	Shipping    Shipping    `gorm:"foreignKey:OrderId"`
	OrderItem   []OrderItem `gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	ID        uint            `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"orderItemId"`
	Quality   uint            `gorm:"type:int(11) NOT NULL;"`
	Price     decimal.Decimal `gorm:"type:decimal(19,4) NOT NULL;"`
	ProductId uint            `gorm:"type:bigint(20) NOT NULL;index:idx_product_id;" json:"productId"`
	OrderId   uint            `gorm:"type:bigint(20) NOT NULL;index:idx_order_id;" json:"orderId"`
	Product   Product         `gorm:"foreignKey:ProductId"`
}

type OrderUseCase interface {
	GetAll(ctx context.Context) ([]Order, error)
	GetById(ctx context.Context, id uint) (*Order, error)
	GetByUserId(ctx context.Context, userId uint) ([]Order, error)
	Create(ctx context.Context, order *dto.AddOrderDTO) (*Order, error)
	Update(ctx context.Context, order *Order) error
}

type OrderRepository interface {
	GetAll(ctx context.Context) ([]Order, error)
	GetById(ctx context.Context, id uint) (*Order, error)
	GetByUserId(ctx context.Context, userId uint) ([]Order, error)
	Create(ctx context.Context, order *Order) error
	Update(ctx context.Context, order *Order) error
	WithTransaction(ctx context.Context, txFunc func(*gorm.DB) error) (err error)
}
