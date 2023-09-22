package repositories

import (
	"core-api/internal/api/models"
	"core-api/pkg/database"
	"core-api/pkg/pagination"

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
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		db: database.SetupDatabaseConnection(),
	}
}

func (ur *UserRepositoryImpl) All() []models.User {
	var users []models.User

	paginationData := pagination.Pagination{
		Limit:      1,
		Page:       2,
		TotalRows:  100,
		TotalPages: 10,
	}

	ur.db.Scopes(pagination.Paginate(users, &paginationData, ur.db)).Find(&users)

	return users
}

func (ur *UserRepositoryImpl) Create(user models.User) error {
	result := ur.db.Create(&user)
	return result.Error
}

func (ur *UserRepositoryImpl) Find(id int) (*models.User, error) {
	var user models.User
	result := ur.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepositoryImpl) Update(id int, user models.User) error {
	result := ur.db.Model(&models.User{}).Where("id = ?", id).Updates(user)
	return result.Error
}

func (ur *UserRepositoryImpl) Delete(id int) error {
	result := ur.db.Delete(&models.User{}, id)
	return result.Error
}
