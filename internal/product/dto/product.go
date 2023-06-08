package dto

import "mime/multipart"

type ProductGetByCategoryNameDTO struct {
	ProductCategory string `json:"categoryName" binding:"required"`
}

type AddProductDTO struct {
	ProductName     string                  `form:"productName" binding:"required"`
	Price           string                  `form:"price" binding:"required"`
	Introduction    string                  `form:"introduction"`
	Information     string                  `form:"information"`
	Inventory       uint                    `form:"inventory" binding:"required"`
	ProductImage    []*multipart.FileHeader `form:"productImage"`
	ProductCategory string                  `form:"categoryName" binding:"required"`
}
