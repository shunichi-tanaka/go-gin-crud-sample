package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	service := NewUserService()

	expect := "1"
	user, err := service.GetUser("1")
	if err != nil {
		return
	}
	result := user.ID
	assert.Equal(t, expect, result)
}
