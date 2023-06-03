package dto

type AddProductCategoryDTO struct {
	CategoryName string `json:"categoryName" binding:"required"`
}

type UpdateProductCategoryDTO struct {
	ID           uint   `json:"productCategoryId" binding:"required"`
	CategoryName string `json:"categoryName" binding:"required"`
}