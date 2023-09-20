package controllers

import (
	"core-api/internal/api/services"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUsers(ctx *gin.Context)
}

type ControllerConnector struct {
	userService services.UserService
}

func NewUserController() UserController {
	return &ControllerConnector{
		userService: services.NewUserService(),
	}
}

func (uc *ControllerConnector) GetAllUsers(ctx *gin.Context) {
	users := uc.userService.GetAllUsers()
	ctx.JSON(200, users)
}
