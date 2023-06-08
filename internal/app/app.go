package app

import (
	"HaveBing-Backend/internal/middleware/error"
	_productCategoryDelivery "HaveBing-Backend/internal/product_category/delivery"
	_productCategoryRepository "HaveBing-Backend/internal/product_category/repository"
	_productCategoryUsecase "HaveBing-Backend/internal/product_category/usecase"

	_userDelivery "HaveBing-Backend/internal/user/delivery"
	_userRepository "HaveBing-Backend/internal/user/repository"
	_userUsecase "HaveBing-Backend/internal/user/usecase"

	_productDelivery "HaveBing-Backend/internal/product/delivery"
	_productRepository "HaveBing-Backend/internal/product/repository"
	_productUsecase "HaveBing-Backend/internal/product/usecase"

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

	userRepository := _userRepository.New(db)
	userUsecase := _userUsecase.New(userRepository)
	_userDelivery.Register(router, userUsecase)

	productRepository := _productRepository.New(db)
	productUsecase := _productUsecase.New(productRepository, productCategoryRepository)
	_productDelivery.Register(router, productUsecase)

	return app
}
