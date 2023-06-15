package domain

import "context"

type Shipping struct {
	ID             uint   `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"shippingId"`
	RecipientName  string `gorm:"type:varchar(255) NOT NULL;" json:"recipientName"`
	RecipientPhone string `gorm:"type:varchar(50) NOT NULL;" json:"recipientPhone"`
	Address        string `gorm:"type:mediumtext NOT NULL;" json:"address"`
	ShippingMethod string `gorm:"type:varchar(255) NOT NULL;" json:"shippingMethod"`
	OrderId        uint   `gorm:"type:bigint(20) NOT NULL;index:idx_order_id;" json:"orderId"`
}

type ShippingRepository interface {
	GetById(ctx context.Context, id uint) (*Shipping, error)
	Create(ctx context.Context, shipping *Shipping) error
	Update(ctx context.Context, shipping *Shipping) error
}
