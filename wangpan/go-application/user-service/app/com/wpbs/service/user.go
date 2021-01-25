package service

import (
	"context"
	"time"
)

import (
	"user-service/app/com/wpbs/pojo"
)

import (
	gxlog "github.com/dubbogo/gost/log"
)

// 用户业务逻辑处理对象
type UserService struct {
}

func (u *UserService) GetUser(ctx context.Context, req []interface{}) (*pojo.User, error) {
	gxlog.CInfo("req:%#v", req)
	rsp := pojo.User{"A001", "Alex Stocks", 18, time.Now()}
	gxlog.CInfo("rsp:%#v", rsp)
	return &rsp, nil
}

// 继承RPCSerive 接
// 口  实现Reference方法:
func (u *UserService) Reference() string {
	return "UserService"
}
