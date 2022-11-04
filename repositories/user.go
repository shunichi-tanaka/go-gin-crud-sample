package repositories

import (
	"go-gin-crud-sample/models"
)

type UserRepository interface {
	FindByID(ID string) (*models.User, error)
}

type userRepository struct {
}

func NewUserRepository() userRepository {
	return &userRepository{}
}

func (ur *userRepository) FindByID(ID string) (*models.User, error) {

	user := models.User{
		ID:   "1",
		Name: "Bob",
	}

	return &user, nil
}
