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
		log.Errorf("traceId:%v param bind error: #v", user.TraceId, err)
		Error(ctx, common.CODE_405)
		return
	}

	log.Info(" traceId:%v controller createUser ", user.TraceId)
	userService = config.GetConsumerService("UserService").(*service.UserService)
	userDTO := userBD(user)
	err = userService.CreateUser(context.TODO(), []interface{}{userDTO}, userDTO)
	if err != nil {
		log.Errorf("traceId:%v  user controller createUser error", user.TraceId, err)
		Error(ctx, common.CODE_500)
		return
	}

	Success(ctx, "注册成功", gin.H{"info": userDTO})
	log.Infof("traceId:%v controller createUser success", user.TraceId)
}

// 通过ID获取用户
func GetUserById(ctx *gin.Context) {
	user := &BO.GetUserById{}
	var err error
	err = ctx.ShouldBind(user)
	if err != nil {
		log.Errorf("traceId:%v param bind error: #v", user.TraceId, err)
		Error(ctx, common.CODE_405)
		return
	}

	user.Id = 1
	user.TraceId = "111"
	log.Infof("traceId:%v controller GetUserById", user.TraceId)
	userService = config.GetConsumerService("UserService").(*service.UserService)
	userDTO := &DTO.User{}
	err = userService.GetUserById(context.TODO(), []interface{}{user.TraceId, user.Id}, userDTO)
	if err != nil {
		log.Errorf("traceId:%v  controller call userService error:%v", err)
		Error(ctx, common.CODE_500)
		return
	}

	Success(ctx, "获取成功", gin.H{"info": userDTO})
	log.Infof("traceId:%v controller getUserById ", user.TraceId)
}

// DTO <-> BO
func userBD(user *BO.User) *DTO.User {
	return &DTO.User{
		TraceId:     user.TraceId,
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
		TraceId:     user.TraceId,
		Id:          user.Id,
		Name:        user.Name,
		PassWord:    user.PassWord,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		CreateTime:  user.CreateTime,
		UpdateTime:  user.UpdateTime,
	}
}
