package middlewares

import (
	"core-api/pkg/helpers"
	"core-api/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")

		if auth == "" {
			response := helpers.BuildErrorResponse("Failed to process request", "No token found", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		parts := strings.Split(auth, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			response := helpers.BuildErrorResponse("Failed to process request", "Invalid token format", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := parts[1]

		jwtService := jwt.NewJWTService()
		claims, err := jwtService.VerifyToken(tokenString)
		if err != nil {
			response := helpers.BuildErrorResponse("Invalid token", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set("claims", claims)
	}
}
