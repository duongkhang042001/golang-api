package repositories

import (
	"core-api/internal/api/configs"
	"core-api/internal/api/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	All() []models.User
	Create(user models.User) error
	Find(id int) (*models.User, error)
	Update(id int, user models.User) error
	Delete(id int) error
}

type UserRepositoryImpl struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		connection: configs.SetupDatabaseConnection(),
	}
}

func (ur *UserRepositoryImpl) All() []models.User {
	var users []models.User
	ur.connection.Find(&users)
	return users
}

func (ur *UserRepositoryImpl) Create(user models.User) error {
	result := ur.connection.Create(&user)
	return result.Error
}

func (ur *UserRepositoryImpl) Find(id int) (*models.User, error) {
	var user models.User
	result := ur.connection.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepositoryImpl) Update(id int, user models.User) error {
	result := ur.connection.Model(&models.User{}).Where("id = ?", id).Updates(user)
	return result.Error
}

func (ur *UserRepositoryImpl) Delete(id int) error {
	result := ur.connection.Delete(&models.User{}, id)
	return result.Error
}
