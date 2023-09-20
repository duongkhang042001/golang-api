package services

import (
	"core-api/internal/api/models"
	"core-api/internal/api/repositories"
)

type UserService interface {
	GetAllUsers() []models.User
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService() UserService {
	return &UserServiceImpl{
		userRepository: repositories.NewUserRepository(),
	}
}

func (us *UserServiceImpl) GetAllUsers() []models.User {
	return us.userRepository.All()
}
