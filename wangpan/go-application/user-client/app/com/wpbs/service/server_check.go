package service

import (
	"context"
	"user-client/app/com/wpbs/pojo"
)

type ServerCheckService struct {
	Check func(ctx context.Context, req []interface{}, rsp *pojo.ServerCheck) error
}

// 	Reference
func (s *ServerCheckService) Reference() string {
	return "ServerCheckService"
}
