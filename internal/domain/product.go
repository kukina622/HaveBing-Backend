package domain

import (
	"context"
	"mime/multipart"

	"github.com/shopspring/decimal"
)

type ProductImage struct {
	ID        uint   `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"productImageId"`
	Path      string `gorm:"type:varchar(255) NOT NULL;"`
	ProductId uint   `gorm:"type:bigint(20) NOT NULL;"`
}

type Product struct {
	ID                uint            `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;"`
	ProductName       string          `gorm:"type:mediumtext NOT NULL;"`
	Price             decimal.Decimal `gorm:"type:decimal(19,4) NOT NULL;"`
	Introduction      string          `gorm:"type:longtext;"`
	Information       string          `gorm:"type:longtext;"`
	Inventory         uint            `gorm:"type:int(11) NOT NULL; check:inventory>=0;"`
	ProductCategoryId uint            `gorm:"type:bigint(20) NOT NULL;"`
	ProductImage      []ProductImage  `gorm:"foreignKey:ProductId"`
	ProductCategory   ProductCategory
}

type ProductRepository interface {
	GetAll(ctx context.Context) ([]Product, error)
	GetById(ctx context.Context, id uint) (*Product, error)
	GetByCategoryId(ctx context.Context, categoryId uint) ([]Product, error)
	Save(ctx context.Context, product *Product) error
}

type ProductUseCase interface {
	GetAll(ctx context.Context) ([]Product, error)
	GetById(ctx context.Context, id uint) (*Product, error)
	GetByCategoryId(ctx context.Context, categoryId uint) ([]Product, error)
	GetByCategoryName(ctx context.Context, categoryName string) ([]Product, error)
	Create(ctx context.Context, product *Product, productImages []*multipart.FileHeader, categoryName string) error
	Update(ctx context.Context, product *Product) error
}
