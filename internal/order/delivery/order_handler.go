package delivery

import (
	"HaveBing-Backend/internal/domain"
	"HaveBing-Backend/internal/middleware/error"
	"net/http"
	"strconv"

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
	router.GET("/user/:userId/order", handler.GetByUserId)
}

func (handler *OrderHandler) GetAll(ctx *gin.Context) {
	orderList, err := handler.orderUsecase.GetAll(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, orderList)
}

func (handler *OrderHandler) GetByUserId(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	orderList, err := handler.orderUsecase.GetByUserId(ctx, uint(userId))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, &error.ServerError{
			Code: http.StatusNotFound,
			Msg:  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, orderList)
}