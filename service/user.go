package service

import (
	"go-gin-crud-sample/models"
)

type UserService interface {
	GetUser(ID string) (*models.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (us *userService) GetUser(ID string) (*models.User, error) {

	user := models.User{
		ID:   ID,
		Name: "Bob",
	}

	return &user, nil
}
