package usecase

import (
	"HaveBing-Backend/internal/domain"
	"context"
)

type ProductCategoryUseCase struct {
	repo domain.ProductCategoryRepository
}

func New(repo domain.ProductCategoryRepository) domain.ProductCategoryUseCase {
	return &ProductCategoryUseCase{
		repo: repo,
	}
}

func (p *ProductCategoryUseCase) GetAll(ctx context.Context) ([]domain.ProductCategory, error) {
	return p.repo.GetAll(ctx)
}

func (p *ProductCategoryUseCase) GetById(ctx context.Context, id int) (*domain.ProductCategory, error) {
	return p.repo.GetById(ctx, id)
}

func (p *ProductCategoryUseCase) Update(ctx context.Context, productCategory *domain.ProductCategory) error {
	return nil
}

func (p *ProductCategoryUseCase) Save(ctx context.Context, productCategory *domain.ProductCategory) error {
	return p.repo.Save(ctx, productCategory)
}
