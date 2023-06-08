package usecase

import (
	"HaveBing-Backend/internal/domain"
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"HaveBing-Backend/internal/util/array"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	return p.productRepo.GetByCategoryId(ctx, categoryId)
}

func (p *ProductUseCase) GetByCategoryName(ctx context.Context, categoryName string) ([]domain.Product, error) {
	productCategory, err := p.productCategoryRepo.GetByName(ctx, categoryName)
	if err != nil {
		return nil, err
	}
	return p.productRepo.GetByCategoryId(ctx, productCategory.ID)
}

func (p *ProductUseCase) Create(ctx context.Context, product *domain.Product, productImages []*multipart.FileHeader, categoryName string) error {
	productCategory, err := p.productCategoryRepo.GetByName(ctx, categoryName)
	if err != nil {
		return err
	}
	product.ProductCategory = *productCategory
	if err := p.productRepo.Save(ctx, product); err != nil {
		return err
	}

	imagePaths, err := p.saveImageWithPrefix(ctx, productImages, fmt.Sprintf("%d-", product.ID))
	if err != nil {
		return err
	}

	productImageList := array.Map(
		imagePaths,
		func(path string) domain.ProductImage {
			return domain.ProductImage{
				ProductId: product.ID,
				Path:      path,
			}
		},
	)
	product.ProductImage = productImageList
	if err := p.productRepo.Save(ctx, product); err != nil {
		return err
	}
	return nil
}

func (p *ProductUseCase) Update(ctx context.Context, product *domain.Product) error {
	return nil
}

func (p *ProductUseCase) saveImageWithPrefix(ctx context.Context, productImages []*multipart.FileHeader, prefix string) ([]string, error) {
	exePath, _ := os.Executable()
	rootPath := filepath.Dir(exePath)
	_ctx := ctx.(*gin.Context)
	paths := []string{}
	for _, productImage := range productImages {
		filename := prefix + uuid.New().String() + filepath.Ext(productImage.Filename)
		path := rootPath + "/assets/images/" + filename
		fmt.Println(path)
		if err := _ctx.SaveUploadedFile(productImage, path); err != nil {
			return nil, err
		}
		paths = append(paths, filename)
	}
	return paths, nil
}
