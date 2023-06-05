package auth

import (
	"HaveBing-Backend/internal/middleware/error"
	jwtUtil "HaveBing-Backend/internal/util/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware(ctx *gin.Context) {
	reqToken := ctx.GetHeader("Authorization")
	payload, err := jwtUtil.ExtractPayload(reqToken)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, &error.ServerError{
			Code: http.StatusUnauthorized,
			Msg:  "Unauthorized",
		})
		return
	}
	payloadRole := payload["role"]
	if payloadRole == nil {
		ctx.AbortWithError(http.StatusUnauthorized, &error.ServerError{
			Code: http.StatusUnauthorized,
			Msg:  "Unauthorized",
		})
		return
	}
	roleName := payloadRole.(map[string]any)["roleName"]
	if roleName == "admin" || roleName == "debugger" {
		ctx.Next()
		return
	}
	ctx.AbortWithError(http.StatusUnauthorized, &error.ServerError{
		Code: http.StatusUnauthorized,
		Msg:  "Unauthorized",
	})
}
