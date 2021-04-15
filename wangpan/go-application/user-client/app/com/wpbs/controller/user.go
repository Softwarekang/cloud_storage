package controller

import (
	"context"
	"github.com/apache/dubbo-go/config"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"user-client/app/com/wpbs/BO"
	"user-client/app/com/wpbs/DTO"
	"user-client/app/com/wpbs/service"
	"user-client/common/httpcode"
	"user-client/common/utils"
)

// 用户注册
func CreateUser(ctx *gin.Context) {
	logger.Info("  controller createUser ")
	var user BO.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		logger.Errorf("param bind error: #v", err)
		Error(ctx, httpcode.CODE_405)
		return
	}

	userService := config.GetConsumerService("UserService").(*service.UserService)
	// 查询用户是否存在
	selectUser := &DTO.User{}
	if err := userService.GetUserByName(context.TODO(), user.Name, selectUser); err != nil {
		logger.Errorf("user controller call userService error:%v", err)
		Error(ctx, httpcode.CODE_500)
		return
	}

	// 存在重复用户名
	if isNil := utils.IsNilString(selectUser.Name); !isNil {
		logger.Info("user controller createUser userName exist")
		Error(ctx, httpcode.CODE_406)
		return
	}

	userDTO := userBD(&user)
	if err := userService.CreateUser(context.TODO(), []interface{}{userDTO}, userDTO); err != nil {
		logger.Errorf("user controller call userService createUser error", err)
		Error(ctx, httpcode.CODE_406)
		return
	}

	logger.Info(" controller createUser success")
	Success(ctx, "注册成功", gin.H{"info": userDTO})
}

// 通过ID获取用户
func GetUserById(ctx *gin.Context) {
	logger.Infof("controller GetUserById")
	userId := ctx.Query("id")
	userService := config.GetConsumerService("UserService").(*service.UserService)
	userDTO := &DTO.User{}
	err := userService.GetUserById(context.TODO(), []interface{}{userId}, userDTO)
	if err != nil {
		logger.Errorf(" controller call userService error:%v", err)
		Error(ctx, httpcode.CODE_500)
		return
	}

	Success(ctx, "获取成功", gin.H{"info": userDTO})
	logger.Infof("controller getUserById ")
}

// 用户登录
func Login(ctx *gin.Context) {
	logger.Info("userController login")
	user := &BO.User{}
	if err := ctx.Bind(user); err != nil {
		logger.Errorf("param bind error: #v", err)
		Error(ctx, httpcode.CODE_405)
		return
	}

	userService := config.GetConsumerService("UserService").(*service.UserService)
	selectUser := &DTO.User{}
	if err := userService.GetUserByName(context.TODO(), user.Name, selectUser); err != nil {
		logger.Errorf("user controller Login rpc call userService error", err)
		Error(ctx, httpcode.CODE_500)
		return
	}

	if isNil := utils.IsNilString(selectUser.Name); isNil {
		logger.Infof("user controller login user:%v is not exist", selectUser.Name)
		Error(ctx, httpcode.CODE_407)
		return
	}

	if selectUser.PassWord != user.PassWord {
		logger.Infof("user controller log passWord not equals :source:%v,target:%v", user.PassWord, selectUser.PassWord)
		Error(ctx, httpcode.CODE_408)
		return
	}

	logger.Infof("user controller login success rsp model:%v", *selectUser)
	Success(ctx, "登录成功", gin.H{"info": selectUser})
}

// 用户头像上传
func UserImage(ctx *gin.Context) {
	fileHeader, err := checkUserImageUpload(ctx)
	if err != nil {
		logger.Error("controller UserImage file param error")
		Error(ctx, httpcode.CODE_405)
		return
	}

	fileViewUrl, err := UploadFileMinio(fileHeader)
	if err != nil {
		logger.Errorf("controller UserImage upload file error:%v", err)
		Error(ctx, httpcode.CODE_500)
		return
	}

	logger.Info("controller UserImage success rsp:%v", fileViewUrl)
	Success(ctx, "上传成功", gin.H{"info": fileViewUrl})
}

func checkUserImageUpload(ctx *gin.Context) (*multipart.FileHeader, error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		logger.Errorf("file param err:", err)
		return nil, err
	}

	return file, nil
}

func UpdateUser(ctx *gin.Context) {
	logger.Infof("controller UpdateUser")
	user := &BO.User{}
	if err := ctx.BindJSON(user); err != nil {
		logger.Errorf("param bind error:%v", err)
		Error(ctx, httpcode.CODE_405)
		return
	}

	userService := config.GetConsumerService("UserService").(*service.UserService)
	userDto :=userBD(user)
	if err := userService.UpdateUser(context.TODO(), userDto); err != nil {
		logger.Errorf("user controller UpdateUser call userService error:%v", err)
		Error(ctx, httpcode.CODE_500)
		return
	}

	logger.Infof("user controller UpdateUser success")
	Success(ctx, "更新成功", nil)
}

// DTO <-> BO
func userBD(user *BO.User) *DTO.User {
	return &DTO.User{
		Id:          user.Id,
		Name:        user.Name,
		PassWord:    user.PassWord,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		HeadImageUrl: user.HeadImageUrl,
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
		HeadImageUrl: user.HeadImageUrl,
		CreateTime:  user.CreateTime,
		UpdateTime:  user.UpdateTime,
	}
}
