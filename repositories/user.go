package repositories

import (
	"database/sql"

	"go-gin-crud-sample/models"
)

type UserRepository interface {
	FindByID(ID string) (*models.User, error)
}

type userRepository struct {
	DB *sql.DB
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
