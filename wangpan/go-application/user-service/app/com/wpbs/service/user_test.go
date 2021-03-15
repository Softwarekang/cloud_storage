package service

import (
	"context"
	"gotest.tools/assert"
	"testing"
	"user-service/app/com/wpbs/DTO"
)

var (
	userService = new(UserService)
)

func TestUserService_CreateUser(t *testing.T) {
	_, err := userService.CreateUser(context.TODO(), []interface{}{&DTO.User{Id: 5, Name: "安康", PassWord: "123"}})
	assert.Equal(t, err, nil)
}
