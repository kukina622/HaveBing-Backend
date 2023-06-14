package delivery

import (
	"HaveBing-Backend/internal/domain"
	"HaveBing-Backend/internal/dto"
	"HaveBing-Backend/internal/middleware/auth"
	"HaveBing-Backend/internal/middleware/error"
	jwtUtil "HaveBing-Backend/internal/util/jwt"
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
	router.GET("/user", auth.JwtAuthMiddleware, handler.GetCurrentUser)
	router.GET("/user/all", auth.JwtAuthMiddleware, auth.AdminAuthMiddleware, handler.GetAll)
	router.PATCH("/user/available", auth.JwtAuthMiddleware, auth.AdminAuthMiddleware, handler.ToggleUserAvailable)
	router.PATCH("/user", auth.JwtAuthMiddleware, handler.Update)
}

func (handler *UserHandler) Login(ctx *gin.Context) {
	var body dto.LoginUserDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	if ok, user, token := handler.userUseCase.Login(ctx, body.Email, body.Password); ok {
		userResponse, _ := dto.NewUserResponse(user).(dto.UserResponseDTO)
		responseBody := &dto.UserWithTokenResponseDTO{
			UserResponseDTO: userResponse,
			Token:           token,
		}
		ctx.JSON(http.StatusOK, responseBody)
		return
	}
	ctx.AbortWithError(http.StatusUnauthorized, &error.ServerError{
		Code: http.StatusUnauthorized,
		Msg:  "Email or password is incorrect",
	})
}

func (handler *UserHandler) Register(ctx *gin.Context) {
	var body dto.RegisterUserDTO
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

func (handler *UserHandler) GetCurrentUser(ctx *gin.Context) {
	user, err := handler.userUseCase.GetCurrentUser(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, &error.ServerError{
			Code: http.StatusNotFound,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.NewUserResponse(user))
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
	var body dto.ToggleUserAvailableDTO
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

func (handler *UserHandler) Update(ctx *gin.Context) {
	var body dto.UpdateUserDTO
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

	reqToken := ctx.GetHeader("Authorization")
	payload, err := jwtUtil.ExtractPayload(reqToken)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	user := domain.User{
		ID:       uint(payload["userId"].(float64)),
		Email:    body.Email,
		Name:     body.Name,
		Birthday: birthday,
		Phone:    body.Phone,
	}

	if err := handler.userUseCase.Update(ctx, &user, body.OldPassword, body.NewPassword); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, &error.ServerError{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	ctx.Status(http.StatusOK)
}
