package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	for _, err := range ctx.Errors {
		if serverError ,ok := err.Err.(*ServerError); ok {
			ctx.JSON(serverError.Code, serverError)
		} else {
			ctx.Status(http.StatusInternalServerError)
		}
	}
}