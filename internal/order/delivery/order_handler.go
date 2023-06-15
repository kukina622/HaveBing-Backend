package delivery

import (
	"HaveBing-Backend/internal/domain"
	"HaveBing-Backend/internal/dto"
	"HaveBing-Backend/internal/dto/request"
	"HaveBing-Backend/internal/dto/response"
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
	router.GET("/order/:id", handler.GetById)
	router.GET("/user/:userId/order", handler.GetByUserId)
	router.POST("/order", handler.Create)
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
	ctx.JSON(http.StatusOK, response.NewOrderResponseDTO(orderList))
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
	ctx.JSON(http.StatusOK, response.NewOrderResponseDTO(orderList))
}

func (handler *OrderHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	order, err := handler.orderUsecase.GetById(ctx, uint(id))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, &error.ServerError{
			Code: http.StatusNotFound,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.NewOrderResponseDTO(order))
}

func (handler *OrderHandler) Create(ctx *gin.Context) {
	var body request.AddOrderRequestDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	productList := []struct {
		ProductId uint
		Quantity  uint
	}{}

	for _, product := range body.ProductList {
		productList = append(productList, struct {
			ProductId uint
			Quantity  uint
		}{
			ProductId: product.ProductId,
			Quantity:  product.Quantity,
		})
	}

	newOder := dto.AddOrderDTO{
		UserId:         body.UserId,
		Note:           body.Note,
		RecipientName:  body.RecipientName,
		RecipientPhone: body.RecipientPhone,
		Address:        body.Address,
		ProductList:    productList,
		ShippingMethod: body.ShippingMethod,
	}
	order, err := handler.orderUsecase.Create(ctx, &newOder)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.NewOrderResponseDTO(order))
}
