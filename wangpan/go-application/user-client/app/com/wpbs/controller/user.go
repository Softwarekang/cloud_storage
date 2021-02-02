package controller

import (
	"context"
	"github.com/apache/dubbo-go/config"
	"github.com/gin-gonic/gin"
	"user-client/app/com/wpbs/BO"
	"user-client/app/com/wpbs/DTO"
	"user-client/app/com/wpbs/service"
	"user-client/common"
)

var (
	userService *service.UserService
)

// 用户注册
func CreateUser(ctx *gin.Context) {
	user := &BO.User{}
	var err error
	err = ctx.ShouldBind(user)
	if err != nil {
		log.Errorf("param bind error: #v", err)
		Error(ctx, common.CODE_405)
		return
	}

	log.Info("  controller createUser ")
	userService = config.GetConsumerService("UserService").(*service.UserService)
	userDTO := userBD(user)
	err = userService.CreateUser(context.TODO(), []interface{}{userDTO}, userDTO)
	if err != nil {
		log.Errorf("user controller createUser error", err)
		Error(ctx, common.CODE_500)
		return
	}

	Success(ctx, "注册成功", gin.H{"info": userDTO})
	log.Infof("traceId:%v controller createUser success")
}

// 通过ID获取用户
func GetUserById(ctx *gin.Context) {
	log.Infof("controller GetUserById")
	userId := ctx.Query("id")
	userService = config.GetConsumerService("UserService").(*service.UserService)
	userDTO := &DTO.User{}
	err := userService.GetUserById(context.TODO(), []interface{}{userId}, userDTO)
	if err != nil {
		log.Errorf(" controller call userService error:%v", err)
		Error(ctx, common.CODE_500)
		return
	}

	Success(ctx, "获取成功", gin.H{"info": userDTO})
	log.Infof("controller getUserById ")
}

// DTO <-> BO
func userBD(user *BO.User) *DTO.User {
	return &DTO.User{
		Id:          user.Id,
		Name:        user.Name,
		PassWord:    user.PassWord,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		CreateTime:  user.CreateTime,
		UpdateTime:  user.UpdateTime,
	}
}

func userDB(user *DTO.User) *BO.User {
	return &BO.User{
		Id:          user.Id,
		Name:        user.Name,
		PassWord:    user.PassWord,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		CreateTime:  user.CreateTime,
		UpdateTime:  user.UpdateTime,
	}
}
