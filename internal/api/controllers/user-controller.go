package controllers

type UserController interface {
}

type ControllerConnector struct {
}

func NewUserController() UserController {
	return &ControllerConnector{}
}
