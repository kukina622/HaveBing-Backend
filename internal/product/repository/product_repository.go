package repository

import (
	"HaveBing-Backend/internal/domain"
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var productList []domain.Product
	err := p.db.Preload(clause.Associations).Find(&productList).Error
	return productList, err
}

func (p *ProductRepository) GetById(ctx context.Context, id uint) (*domain.Product, error) {
	var product domain.Product
	err := p.db.Preload(clause.Associations).First(&product, id).Error
	return &product, err
}

func (p *ProductRepository) GetByCategoryId(ctx context.Context, categoryId uint) ([]domain.Product, error) {
	var productList []domain.Product
	err := p.db.Preload(clause.Associations).Where("product_category_id = ?", categoryId).Find(&productList).Error
	return productList, err
}

func (p *ProductRepository) Save(ctx context.Context, product *domain.Product) error {
	return p.db.Save(product).Error
}

func (p *ProductRepository) DecreaseInventoryWithTx(ctx context.Context, tx *gorm.DB, id uint, quantity uint) error {
	result := tx.Model(&domain.Product{ID: id}).Where(
		"inventory - ? >= 0", quantity,
	).Update(
		"inventory", gorm.Expr("inventory - ?", quantity),
	)

	if result.RowsAffected == 0 {
		return fmt.Errorf("inventory not enough")
	}

	return result.Error
}

func (p *ProductRepository) GetByIdWithTx(ctx context.Context, tx *gorm.DB, id uint) (*domain.Product, error) {
	var product domain.Product
	err := tx.Preload(clause.Associations).First(&product, id).Error
	return &product, err
}
