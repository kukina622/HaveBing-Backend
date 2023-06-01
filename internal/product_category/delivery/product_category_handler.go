package delivery

import (
	"HaveBing-Backend/internal/domain"
	"net/http"

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
	router.POST("/productCategory", handler.PostNewProductCategory)
}

func (handler *ProductCategoryHandler) GetAll(ctx *gin.Context) {
	productCategoryList, err := handler.productCategoryUseCase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"productCategory": productCategoryList,
	})
}

func (handler *ProductCategoryHandler) PostNewProductCategory(ctx *gin.Context) {
	var body domain.ProductCategory
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := handler.productCategoryUseCase.Save(ctx, &body); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusOK)
}
