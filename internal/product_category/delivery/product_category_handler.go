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
