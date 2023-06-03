package delivery

import (
	"HaveBing-Backend/internal/domain"
	"HaveBing-Backend/internal/middleware/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func Register(router *gin.RouterGroup, userUsecase domain.UserUseCase) {
	handler := &UserHandler{
		userUseCase: userUsecase,
	}
	router.POST("/login", handler.Login)
	router.POST("/register", handler.Register)
}

func (handler *UserHandler) Login(ctx *gin.Context) {
	var body domain.UserLogin
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	if user, ok := handler.userUseCase.Login(ctx, body.Email, body.Password); ok {
		ctx.JSON(http.StatusOK, user)
		return
	}
	ctx.AbortWithError(http.StatusUnauthorized, &error.ServerError{
		Code: http.StatusUnauthorized,
		Msg:  "Email or password is incorrect",
	})
}

func (handler *UserHandler) Register(ctx *gin.Context) {
	var body domain.User
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	if err := handler.userUseCase.Register(ctx, &body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}
