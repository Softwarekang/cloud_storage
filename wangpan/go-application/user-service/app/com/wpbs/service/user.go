package service

import (
	"context"
	"time"
	"user-service/app/com/wpbs/dao"
	"user-service/app/com/wpbs/store"
	log2 "user-service/common/log"
)

import (
	"user-service/app/com/wpbs/DTO"
)

var (
	log     = log2.GetLogger()
	userDao = new(dao.UserDao)
)

// 用户业务逻辑处理对象
type UserService struct {
}

// 模拟获取用户
func (u *UserService) GetUser(ctx context.Context, req []interface{}) (*DTO.User, error) {
	log.Info("GetUser req:", req)
	rsp := DTO.User{1, "Alex Stocks", "rz@ak10.16", "13759972100", "jingyechenfu@aliyun.com",
		time.Now().Unix(), time.Now().Unix()}
	log.Info("rsp:", rsp)
	return &rsp, nil
}

// 注册用户
func (u *UserService) CreateUser(ctx context.Context, req []interface{}) (*DTO.User, error) {
	user := req[0].(*DTO.User)
	log.Infof("UserService createUser req:%v", user)
	engine := store.DBClient.Begin()
	insertedID, err := store.DBClient.User(engine).CreateUser(user)
	if err != nil {
		log.Errorf(" UserService CreateUser error info :%v", err)
		return nil, err
	}

	user.Id = insertedID
	log.Infof(" UserService createUser success")
	return user, nil
}

// 获取用户
func (u *UserService) GetUserById(ctx context.Context, req []interface{}) (*DTO.User, error) {
	userId := req[0].(string)
	log.Infof(" UserService GetUserById")
	engine := store.DBClient.Begin()
	user, err := store.DBClient.User(engine).GetUserById(userId)
	if err != nil {
		log.Errorf("UserService GetUserById error info:%v ", err)
		return nil, err
	}

	log.Info(" UserService GetUserById success")

	return user, nil
}

func (u *UserService) GetUserByName(ctx context.Context, req interface{}) (*DTO.User, error) {
	log.Info("UserService GetUserByName")
	name := req.(string)
	engine := store.DBClient.Begin()
	user, err := store.DBClient.User(engine).GetUserByName(name)
	if err != nil {
		log.Errorf("UserService GetUserByName error info:%v ", err)
		return nil, err
	}

	log.Info("UserService GetUserByName success")
	return user, nil
}

// 继承RPCSerive 接
// 口  实现Reference方法:
func (u *UserService) Reference() string {
	return "UserService"
}
