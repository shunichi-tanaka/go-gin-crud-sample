package service

import (
	"go-gin-crud-sample/models"
)

type UserService interface {
	GetUser(ID string) (*models.User, error)
	ListUser() ([]*models.User, error)
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

func (us *userService) ListUser() ([]*models.User, error) {

	user1 := models.User{
		ID:   ID1,
		Name: "Bob",
	}

	user2 := models.User{
		ID:   ID2,
		Name: "Alice",
	}
	
	ver users []*models.User
	users = append(users, &user1)
	users = append(users, &user2)

	return users, nil
}
