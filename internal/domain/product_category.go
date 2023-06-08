package domain

import "context"

type ProductCategory struct {
	ID           uint   `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"productCategoryId"`
	CategoryName string `gorm:"type:varchar(50) NOT NULL;" json:"categoryName"`
}

type ProductCategoryRepository interface {
	GetAll(ctx context.Context) ([]ProductCategory, error)
	GetById(ctx context.Context, id int) (*ProductCategory, error)
	GetByName(ctx context.Context, productCategoryName string) (*ProductCategory, error)
	Update(ctx context.Context, productCategory *ProductCategory) error
	Save(ctx context.Context, productCategory *ProductCategory) error
}

type ProductCategoryUseCase interface {
	GetAll(ctx context.Context) ([]ProductCategory, error)
	GetById(ctx context.Context, id int) (*ProductCategory, error)
	Update(ctx context.Context, productCategory *ProductCategory) error
	Save(ctx context.Context, productCategory *ProductCategory) error
}
