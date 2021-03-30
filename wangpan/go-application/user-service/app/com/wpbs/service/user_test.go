package service

import (
	"gotest.tools/assert"
	"path"
	"testing"
)

var (
	userService = new(UserService)
)

func TestUserService_CreateUser(t *testing.T) {
	/*_, err := userService.CreateUser(context.TODO(), []interface{}{&DTO.User{Id: 5, Name: "安康", PassWord: "123"}})
	assert.Equal(t, err, nil)*/
	assert.Equal(t,getSuffix("user.jpg"),"jpg")
}

// 获取文件后缀
func getSuffix(fileName string) string {
	suffix := path.Ext(fileName)
	if suffix == "" {
		suffix = ".default"
	}
	return suffix[1:]
}