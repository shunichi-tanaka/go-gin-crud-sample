package repositories

import (
	"gorm.io/gorm"

	"go-gin-crud-sample/models"
)

type UserRepository interface {
	FindByID(id string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) FindByID(id string) (*models.User, error) {

	user := models.User{}
	err := ur.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
