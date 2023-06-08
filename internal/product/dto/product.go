package dto

type ProductGetByCategoryNameDTO struct {
	ProductCategory string `json:"categoryName" binding:"required"`
}
