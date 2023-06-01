package app

import (
	"HaveBing-Backend/internal/middleware/error"
	_productCategoryDelivery "HaveBing-Backend/internal/product_category/delivery"
	_productCategoryRepository "HaveBing-Backend/internal/product_category/repository"
	_productCategoryUsecase "HaveBing-Backend/internal/product_category/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitApplication(db *gorm.DB) *gin.Engine {
	app := gin.Default()

	app.Use(error.ErrorHandler)

	router := app.Group("/api")

	productCategoryRepository := _productCategoryRepository.New(db)
	productCategoryUsecase := _productCategoryUsecase.New(productCategoryRepository)
	_productCategoryDelivery.Register(router, productCategoryUsecase)

	return app
}
