package delivery

import (
	"HaveBing-Backend/internal/domain"
	"HaveBing-Backend/internal/middleware/error"
	"HaveBing-Backend/internal/product/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUseCase domain.ProductUseCase
}

func Register(router *gin.RouterGroup, productUseCase domain.ProductUseCase) {
	handler := &ProductHandler{
		productUseCase: productUseCase,
	}
	router.GET("/product", handler.GetAllProduct)
}

func (handler *ProductHandler) GetAllProduct(ctx *gin.Context) {
	product, err := handler.productUseCase.GetAll(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, &error.ServerError{
			Code: http.StatusNotFound,
			Msg:  err.Error(),
		})
		return
	}
	res := dto.NewProductResponse(product)
	ctx.JSON(http.StatusOK, res)
}
