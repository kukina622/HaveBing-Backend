package delivery

import (
	"HaveBing-Backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUsecase domain.OrderUseCase
}

func Register(router *gin.RouterGroup, orderUsecase domain.OrderUseCase) {
	handler := &OrderHandler{
		orderUsecase: orderUsecase,
	}
	router.GET("/order", handler.GetAll)
}

func (handler *OrderHandler) GetAll(ctx *gin.Context) {
}