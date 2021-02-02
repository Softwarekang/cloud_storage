package dao

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
	"time"
	"user-service/app/com/wpbs/DTO"
)

var (
	userDao = new(UserDao)
)

/// test getUserByID
func TestUserDao_GetUserById(t *testing.T) {
	user, _ := userDao.GetUserById("1")
	assert.Equal(t, user.Id, int64(1))
}

// test createUser
func TestUserDao_CreateUser(t *testing.T) {
	user := &DTO.User{
		Name:        "安康",
		PhoneNumber: "13759972100",
		PassWord:    "13759972100",
		Email:       "jingyechenfu@aliyun.com",
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	}

	insertedID, e := userDao.CreateUser(user)
	fmt.Println(insertedID)
	err := e
	assert.Equal(t, err, nil)
}

// test getUserByName
func TestUserDao_GetUserByName(t *testing.T) {
	user, err := userDao.GetUserByName("安康")
	fmt.Println(user)
	assert.Equal(t, err,nil)
}