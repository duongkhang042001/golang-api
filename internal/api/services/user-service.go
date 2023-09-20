package services

import (
	"core-api/internal/api/models"
	"core-api/internal/api/repositories"
)

type UserService interface {
	GetAllUsers() []models.User
}

type ServiceConnector struct {
	UserRepository repositories.UserRepository
}

func NewUserService() UserService {
	return &ServiceConnector{
		UserRepository: repositories.NewUserRepository(),
	}
}

func (service *ServiceConnector) GetAllUsers() []models.User {
	return service.UserRepository.All()
}
