package routes

import (
	"core-api/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	router.GET("users", controllers.NewUserController().GetAllUsers)
}
