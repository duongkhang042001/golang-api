package middlewares

import (
	"core-api/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Missing authorization header"})
			ctx.Abort()
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
			ctx.Abort()
			return
		}

		tokenString := parts[1]

		jwtService := jwt.NewJWTService()

		claims, err := jwtService.VerifyToken(tokenString)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
	}
}
