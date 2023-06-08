package dto

import (
	"HaveBing-Backend/internal/domain"

	"github.com/shopspring/decimal"
)

type ProductResponseDTO struct {
	ID              uint            `json:"productId"`
	ProductName     string          `json:"productName"`
	Price           decimal.Decimal `json:"price"`
	Introduction    string          `json:"introduction"`
	Information     string          `json:"information"`
	Inventory       uint            `json:"inventory"`
	ProductImage    []string        `json:"productImage"`
	ProductCategory string          `json:"categoryName"`
}

func NewProductResponse[in *domain.Product | []domain.Product](product in) any {
	switch p := any(product).(type) {
	case *domain.Product:
		return ProductResponseDTO{
			ID:              p.ID,
			ProductName:     p.ProductName,
			Price:           p.Price,
			Introduction:    p.Introduction,
			Information:     p.Information,
			Inventory:       p.Inventory,
			ProductImage:    getProductImagePath(p.ProductImage),
			ProductCategory: p.ProductCategory.CategoryName,
		}
	case []domain.Product:
		result := []ProductResponseDTO{}
		for _, _product := range p {
			result = append(result, ProductResponseDTO{
				ID:              _product.ID,
				ProductName:     _product.ProductName,
				Price:           _product.Price,
				Introduction:    _product.Introduction,
				Information:     _product.Information,
				Inventory:       _product.Inventory,
				ProductImage:    getProductImagePath(_product.ProductImage),
				ProductCategory: _product.ProductCategory.CategoryName,
			})
		}
		return result
	}
	return nil
}

func getProductImagePath(productImage []domain.ProductImage) []string {
	result := []string{}
	for _, image := range productImage {
		result = append(result, image.Path)
	}
	return result
}
