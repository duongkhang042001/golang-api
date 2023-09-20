package repositories

import (
	"core-api/internal/api/configs"
	"core-api/internal/api/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	All() []models.User
	Create(user models.User) models.User
	Find(id int) models.User
	Update(id int, user models.User) models.User
	Delete(id int) bool
}

type RepositoryConnector struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &RepositoryConnector{
		connection: configs.SetupDatabaseConnection(),
	}
}

func (rc *RepositoryConnector) All() []models.User {
	var users []models.User
	rc.connection.Find(&users)
	return users
}

func (rc *RepositoryConnector) Create(user models.User) models.User {
	return models.User{}
}

func (rc *RepositoryConnector) Find(id int) models.User {
	return models.User{}
}

func (rc *RepositoryConnector) Update(id int, user models.User) models.User {
	return models.User{}
}

func (rc *RepositoryConnector) Delete(id int) bool {
	return true
}