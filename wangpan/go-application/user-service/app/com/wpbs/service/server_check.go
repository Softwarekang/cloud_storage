package service

import (
	"context"
	gxlog "github.com/dubbogo/gost/log"
	"user-service/app/com/wpbs/pojo"
)

// 服务探测接口
type ServerCheckService struct {
}

func (s *ServerCheckService) Check(ctx context.Context, req []interface{}) (*pojo.ServerCheck, error) {
	gxlog.CInfo("req:%#v", req)
	rsp := &pojo.ServerCheck{Code: 200, Message: "success"}
	gxlog.CInfo("rsp:%#v", rsp)

	return rsp, nil
}

// 实现Reference方法  继承RPCService
func (s *ServerCheckService) Reference() string {
	return "ServerCheckService"
}
