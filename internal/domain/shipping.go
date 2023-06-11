package domain

import "context"

type Shipping struct {
	ID             uint   `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"shippingId"`
	RecipientName  string `gorm:"type:varchar(255) NOT NULL;"`
	RecipientPhone string `gorm:"type:varchar(50) NOT NULL;"`
	Address        string `gorm:"type:mediumtext NOT NULL;"`
	ShippingMethod string `gorm:"type:varchar(255) NOT NULL;"`
	OrderId        uint   `gorm:"type:bigint(20) NOT NULL;index:idx_order_id;" json:"orderId"`
}

type ShippingRepository interface {
	GetById(ctx context.Context, id uint) (*Shipping, error)
	Create(ctx context.Context, shipping *Shipping) error
	Update(ctx context.Context, shipping *Shipping) error
}
