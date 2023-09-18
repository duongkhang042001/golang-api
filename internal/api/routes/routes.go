package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	RegisterAuthRoutes(app)
	RegisterUserRoutes(app)
}
