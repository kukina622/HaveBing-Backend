package request

type AddProductCategoryRequestDTO struct {
	CategoryName string `json:"categoryName" binding:"required"`
}

type UpdateProductCategoryRequestDTO struct {
	ID           uint   `json:"productCategoryId" binding:"required"`
	CategoryName string `json:"categoryName" binding:"required"`
}
