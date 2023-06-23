package delivery

import (
	"HaveBing-Backend/internal/domain"
	"HaveBing-Backend/internal/dto/request"
	"HaveBing-Backend/internal/middleware/error"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductCategoryHandler struct {
	productCategoryUseCase domain.ProductCategoryUseCase
}

func Register(router *gin.RouterGroup, productCategoryUseCase domain.ProductCategoryUseCase) {
	handler := &ProductCategoryHandler{
		productCategoryUseCase: productCategoryUseCase,
	}
	router.GET("/productCategory", handler.GetAll)
	router.GET("/productCategory/:id", handler.GetById)
	router.POST("/productCategory", handler.Save)
	router.PATCH("/productCategory", handler.Update)
}

func (handler *ProductCategoryHandler) GetAll(ctx *gin.Context) {
	productCategoryList, err := handler.productCategoryUseCase.GetAll(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, &error.ServerError{
			Code: http.StatusNotFound,
			Msg:  error.RESOURCE_NOT_FOUND,
		})
		return
	}

	ctx.JSON(http.StatusOK, productCategoryList)
}

func (handler *ProductCategoryHandler) Save(ctx *gin.Context) {
	var body request.AddProductCategoryRequestDTO
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  error.MISSING_PARAMETER,
		})
		return
	}

	p := domain.ProductCategory{
		CategoryName: body.CategoryName,
	}

	if err := handler.productCategoryUseCase.Save(ctx, &p); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  error.CREATE_RESOURCE_FAILED,
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (handler *ProductCategoryHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  error.MISSING_PARAMETER,
		})
		return
	}

	productCategory, err := handler.productCategoryUseCase.GetById(ctx, id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  error.RESOURCE_NOT_FOUND,
		})
		return
	}

	ctx.JSON(http.StatusOK, productCategory)
}

func (handler *ProductCategoryHandler) Update(ctx *gin.Context) {
	var body request.UpdateProductCategoryRequestDTO
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  error.MISSING_PARAMETER,
		})
		return
	}
	p := domain.ProductCategory{
		ID:           body.ID,
		CategoryName: body.CategoryName,
	}

	if err := handler.productCategoryUseCase.Update(ctx, &p); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  error.UPDATE_RESOURCE_FAILED,
		})
		return
	}
	ctx.Status(http.StatusOK)
}
