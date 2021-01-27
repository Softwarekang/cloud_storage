package service

import (
	"context"
	"user-client/app/com/wpbs/DTO"
)

type ServerCheckService struct {
	Check func(ctx context.Context, req []interface{}, rsp *DTO.ServerCheck) error
}

// 	Reference
func (s *ServerCheckService) Reference() string {
	return "ServerCheckService"
}
