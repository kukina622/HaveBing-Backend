package delivery

import (
	"HaveBing-Backend/internal/domain"
	"HaveBing-Backend/internal/middleware/error"
	"HaveBing-Backend/internal/product/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUseCase domain.ProductUseCase
}

func Register(router *gin.RouterGroup, productUseCase domain.ProductUseCase) {
	handler := &ProductHandler{
		productUseCase: productUseCase,
	}
	router.GET("/product", handler.GetAll)
	router.GET("/product/category/:categoryid", handler.GetByCategoryId)
	router.POST("product/category", handler.GetByCategoryName)
	router.GET("/product/:id", handler.GetById)
}

func (handler *ProductHandler) GetAll(ctx *gin.Context) {
	product, err := handler.productUseCase.GetAll(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, &error.ServerError{
			Code: http.StatusNotFound,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.NewProductResponse(product))
}

func (handler *ProductHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	product, err := handler.productUseCase.GetById(ctx, uint(id))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, &error.ServerError{
			Code: http.StatusNotFound,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.NewProductResponse(product))
}

func (handler *ProductHandler) GetByCategoryId(ctx *gin.Context) {
	categoryidId, err := strconv.Atoi(ctx.Param("categoryid"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	product, err := handler.productUseCase.GetByCategoryId(ctx, uint(categoryidId))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, &error.ServerError{
			Code: http.StatusNotFound,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.NewProductResponse(product))
}

func (handler *ProductHandler) GetByCategoryName(ctx *gin.Context) {
	var body dto.ProductGetByCategoryNameDTO
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	product, err := handler.productUseCase.GetByCategoryName(ctx, body.ProductCategory)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, &error.ServerError{
			Code: http.StatusNotFound,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.NewProductResponse(product))
}
