package domain

import "context"

type ProductCategory struct {
	ID           uint
	CategoryName string
}

type ProductCategoryRepository interface {
	GetAll(ctx context.Context) ([]ProductCategory, error)
	GetById(ctx context.Context, id int) (*ProductCategory, error)
	Update(ctx context.Context, productCategory *ProductCategory) error
	Save(ctx context.Context, productCategory *ProductCategory) error
}

type ProductCategoryUseCase interface {
	GetAll(ctx context.Context) ([]ProductCategory, error)
	GetById(ctx context.Context, id int) (*ProductCategory, error)
	Update(ctx context.Context, productCategory *ProductCategory) error
	Save(ctx context.Context, productCategory *ProductCategory) error
}
