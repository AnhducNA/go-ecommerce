package middlewares

import (
	"github.com/AnhducNA/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "secret" {
			response.ErrorResponse(ctx, response.ErrorUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
