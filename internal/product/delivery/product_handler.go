package delivery

import (
	"HaveBing-Backend/internal/domain"
	"HaveBing-Backend/internal/dto"
	"HaveBing-Backend/internal/middleware/auth"
	"HaveBing-Backend/internal/middleware/error"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
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
	router.POST("/product", auth.JwtAuthMiddleware, auth.AdminAuthMiddleware, handler.Save)
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
	var body dto.GetProductByCategoryNameRequestDTO
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

func (handler *ProductHandler) Save(ctx *gin.Context) {
	var body dto.AddProductRequestDTO
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	price, err := decimal.NewFromString(body.Price)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	product := &domain.Product{
		ProductName:  body.ProductName,
		Price:        price,
		Introduction: body.Introduction,
		Information:  body.Information,
		Inventory:    body.Inventory,
	}
	if err := handler.productUseCase.Create(ctx, product, body.ProductImage, body.ProductCategory); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}
