package delivery

import (
	"HaveBing-Backend/internal/domain"
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
		ctx.Error(&error.ServerError{
			Code: http.StatusNotFound,
			Msg: err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, productCategoryList)
}

func (handler *ProductCategoryHandler) Save(ctx *gin.Context) {
	var body domain.ProductCategory
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Error(&error.ServerError{
			Code: http.StatusBadRequest,
			Msg: err.Error(),
		})
		ctx.Abort()
		return
	}

	if err := handler.productCategoryUseCase.Save(ctx, &body); err != nil {
		ctx.Error(&error.ServerError{
			Code: http.StatusBadRequest,
			Msg: err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.Status(http.StatusOK)
}

func (handler *ProductCategoryHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(&error.ServerError{
			Code: http.StatusBadRequest,
			Msg: err.Error(),
		})
		ctx.Abort()
		return
	}

	productCategory, err := handler.productCategoryUseCase.GetById(ctx, id)
	if err != nil {
		ctx.Error(&error.ServerError{
			Code: http.StatusBadRequest,
			Msg: err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, productCategory)
}

func (handler *ProductCategoryHandler) Update(ctx *gin.Context) {
	var body domain.ProductCategory
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Error(&error.ServerError{
			Code: http.StatusBadRequest,
			Msg: err.Error(),
		})
		ctx.Abort()
		return
	}
	if body.ID == 0 {
		ctx.Error(&error.ServerError{
			Code: http.StatusBadRequest,
			Msg: "Missing id field",
		})
		ctx.Abort()
		return
	}
	if err := handler.productCategoryUseCase.Update(ctx, &body); err != nil {
		ctx.Error(&error.ServerError{
			Code: http.StatusBadRequest,
			Msg: err.Error(),
		})
		ctx.Abort()
		return
	}
	ctx.Status(http.StatusOK)
}
