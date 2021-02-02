package service

import (
	"context"
	"fmt"
	gxlog "github.com/dubbogo/gost/log"
	"user-service/app/com/wpbs/DTO"
)

// 服务探测接口
type ServerCheckService struct {
}

func (s *ServerCheckService) Check(ctx context.Context, req []interface{}) (*DTO.ServerCheck, error) {
	gxlog.CInfo("req:%v", req)
	rsp := &DTO.ServerCheck{Code: 200, Message: "success"}
	gxlog.CInfo("rsp:%v", rsp)
	fmt.Println("call success")
	return rsp, nil
}

// 实现Reference方法  继承RPCService
func (s *ServerCheckService) Reference() string {
	return "ServerCheckService"
}
