package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.Next()

	for _, err := range ctx.Errors {
		if serverError ,ok := err.Err.(*ServerError); ok {
			ctx.JSON(serverError.Code, serverError)
		} else {
			ctx.Status(http.StatusInternalServerError)
		}
	}
}