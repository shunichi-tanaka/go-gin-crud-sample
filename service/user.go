package service

import (
	"go-gin-crud-example/models"
)

type UserService interface {
	GetUser(ID string) (*models.User, error)
}

type userService struct {
}

func NewUserService() userService {
	return &userService{}
}

func (us *userService) GetUser(ID string) (*models.User, error) {
	return &models.User{}
}
