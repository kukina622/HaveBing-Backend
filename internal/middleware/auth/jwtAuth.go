package auth

import (
	"HaveBing-Backend/internal/middleware/error"
	jwtUtil "HaveBing-Backend/internal/util/jwt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func JwtAuthMiddleware(ctx *gin.Context) {
	key := os.Getenv("SECRET_KEY")
	reqToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) < 2 {
		ctx.AbortWithError(http.StatusUnauthorized, &error.ServerError{
			Code: http.StatusUnauthorized,
			Msg:  "Unauthorized",
		})
		return
	}
	token := splitToken[1]
	if _, err := jwtUtil.Verify(token, key); err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, &error.ServerError{
			Code: http.StatusUnauthorized,
			Msg:  err.Error(),
		})
		return
	}
	ctx.Next()
}
