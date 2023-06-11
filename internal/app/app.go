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

	_orderDelivery "HaveBing-Backend/internal/order/delivery"
	_orderRepository "HaveBing-Backend/internal/order/repository"
	_orderUsecase "HaveBing-Backend/internal/order/usecase"

	_paymentRepository "HaveBing-Backend/internal/payment/repository"
	_shippingRepository "HaveBing-Backend/internal/shipping/repository"

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

	orderRepository := _orderRepository.New(db)
	paymentRepository := _paymentRepository.New(db)
	shippingRepository := _shippingRepository.New(db)
	orderUsecase := _orderUsecase.New(orderRepository, paymentRepository, shippingRepository)
	_orderDelivery.Register(router, orderUsecase)

	return app
}
