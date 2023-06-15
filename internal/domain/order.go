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
	OrderNumber string      `gorm:"type:varchar(255) NOT NULL;" json:"orderNumber"`
	Status      OrderStatus `gorm:"type:ENUM('preparing', 'shipping', 'done', 'canceled') NOT NULL;" json:"status"`
	UserId      uint        `gorm:"type:bigint(20) NOT NULL;index:idx_user_id;" json:"userId"`
	Note        string      `gorm:"type:longtext" json:"note"`
	OrderDate   time.Time   `gorm:"type:datetime" json:"orderDate"`
	User        User        `gorm:"foreignKey:UserId" json:"user"`
	Payment     Payment     `gorm:"foreignKey:OrderId" json:"payment"`
	Shipping    Shipping    `gorm:"foreignKey:OrderId" json:"shipping"`
	OrderItem   []OrderItem `gorm:"foreignKey:OrderId" json:"orderItem"`
}

type OrderItem struct {
	ID        uint            `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"orderItemId"`
	Quality   uint            `gorm:"type:int(11) NOT NULL;" json:"quality"`
	Price     decimal.Decimal `gorm:"type:decimal(19,4) NOT NULL;" json:"price"`
	ProductId uint            `gorm:"type:bigint(20) NOT NULL;index:idx_product_id;" json:"productId"`
	OrderId   uint            `gorm:"type:bigint(20) NOT NULL;index:idx_order_id;" json:"orderId"`
	Product   Product         `gorm:"foreignKey:ProductId" json:"product"`
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
	CreateWithTx(ctx context.Context, tx *gorm.DB, order *Order) error
}
