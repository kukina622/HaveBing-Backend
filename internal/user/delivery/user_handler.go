package delivery

import (
	"HaveBing-Backend/internal/domain"
	"HaveBing-Backend/internal/middleware/error"
	"HaveBing-Backend/internal/user/dto"
	"net/http"
	"time"

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
	router.GET("/user/all", handler.GetAll)
	router.PATCH("/user/available", handler.ToggleUserAvailable)
}

func (handler *UserHandler) Login(ctx *gin.Context) {
	var body dto.UserLoginDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	if user, ok := handler.userUseCase.Login(ctx, body.Email, body.Password); ok {
		ctx.JSON(http.StatusOK, dto.NewUserResponse(user))
		return
	}
	ctx.AbortWithError(http.StatusUnauthorized, &error.ServerError{
		Code: http.StatusUnauthorized,
		Msg:  "Email or password is incorrect",
	})
}

func (handler *UserHandler) Register(ctx *gin.Context) {
	var body dto.UserRegisterDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	birthday, err := time.Parse("2006-01-02", body.Birthday)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	user := domain.User{
		Email:    body.Email,
		Password: body.Password,
		Name:     body.Name,
		Birthday: birthday,
		Phone:    body.Phone,
	}

	if err := handler.userUseCase.Register(ctx, &user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}

func (handler *UserHandler) GetAll(ctx *gin.Context) {
	users, err := handler.userUseCase.GetAll(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, &error.ServerError{
			Code: http.StatusNotFound,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.NewUserResponse(users))
}

func (handler *UserHandler) ToggleUserAvailable(ctx *gin.Context) {
	var body dto.UserAvailableDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	user := domain.User{
		ID:        body.ID,
		Available: body.Available,
	}

	if err := handler.userUseCase.ToggleUserAvailable(ctx, &user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)
}
