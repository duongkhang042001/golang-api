package controllers

import (
	"core-api/internal/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUsers(ctx *gin.Context)
}

type UserControllerImpl struct {
	userService services.UserService
}

func NewUserController() UserController {
	return &UserControllerImpl{
		userService: services.NewUserService(),
	}
}

func (uc *UserControllerImpl) GetAllUsers(ctx *gin.Context) {
	users := uc.userService.GetAllUsers()
	ctx.JSON(http.StatusOK, users)
}
