package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := uuid.New()
		ctx.Writer.Header().Set("X-Request-Id", uuid.String())
		ctx.Next()
	}
}
