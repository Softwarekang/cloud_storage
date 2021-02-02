package service

import (
	"context"
	"time"
	"user-service/app/com/wpbs/dao"
	"user-service/common"
)

import (
	"user-service/app/com/wpbs/DTO"
)

var (
	log     = common.GetLogger()
	userDao = new(dao.UserDao)
)

// 用户业务逻辑处理对象
type UserService struct {
}

// 模拟获取用户
func (u *UserService) GetUser(ctx context.Context, req []interface{}) (*DTO.User, error) {
	log.Info("GetUser req:", req)
	rsp := DTO.User{"456", 1, "Alex Stocks", "rz@ak10.16", "13759972100", "jingyechenfu@aliyun.com",
		time.Now().Unix(), time.Now().Unix()}
	log.Info("rsp:", rsp)
	return &rsp, nil
}

// 注册用户
func (u *UserService) CreateUser(ctx context.Context, req []interface{}) (*DTO.User, error) {
	user := req[0].(*DTO.User)
	log.Infof("traceId:%v UserService createUser req:", user.TraceId, user)
	insertedID, err := userDao.CreateUser(user)
	if err != nil {
		log.Errorf("traceId:%v UserService CreateUser error info :%v", user.TraceId, err)
		return nil, err
	}

	user.Id = insertedID
	log.Infof("traceId:%v UserService createUser success", user.TraceId)
	return user, nil
}

// 获取用户
func (u *UserService) GetUserById(ctx context.Context, req []interface{}) (*DTO.User, error) {
	traceId, userId := req[0].(string), req[1].(int64)
	log.Infof("traceId:%v UserService GetUserById", traceId)
	user, err := userDao.GetUserById(userId, traceId)
	if err != nil {
		log.Errorf("traceId:%v UserService GetUserById error info:%v ", err)
		return nil, err
	}

	log.Info("traceId:%v UserService GetUserById success", traceId)

	return user, nil
}

// 继承RPCSerive 接
// 口  实现Reference方法:
func (u *UserService) Reference() string {
	return "UserService"
}
