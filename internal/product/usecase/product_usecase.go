package usecase

import (
	"HaveBing-Backend/internal/domain"
	"context"
)

type ProductUseCase struct {
	productRepo         domain.ProductRepository
	productCategoryRepo domain.ProductCategoryRepository
}

func New(productRepo domain.ProductRepository, productCategoryRepo domain.ProductCategoryRepository) domain.ProductUseCase {
	return &ProductUseCase{
		productRepo:         productRepo,
		productCategoryRepo: productCategoryRepo,
	}
}

func (p *ProductUseCase) GetAll(ctx context.Context) ([]domain.Product, error) {
	return p.productRepo.GetAll(ctx)
}

func (p *ProductUseCase) GetById(ctx context.Context, id uint) (*domain.Product, error) {
	return p.productRepo.GetById(ctx, id)
}

func (p *ProductUseCase) GetByCategoryId(ctx context.Context, categoryId uint) ([]domain.Product, error) {
	return nil, nil
}

func (p *ProductUseCase) Create(ctx context.Context, product *domain.Product) error {
	return nil
}

func (p *ProductUseCase) Update(ctx context.Context, product *domain.Product) error {
	return nil

}
