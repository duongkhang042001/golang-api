package routes

import (
	"core-api/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

var userController controllers.UserController = controllers.NewUserController()

func RegisterUserRoutes(router *gin.Engine) {
	router.GET("users", userController.GetAllUsers)
}
