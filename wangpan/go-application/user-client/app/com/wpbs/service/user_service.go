package service

import "context"

import (
	"user-client/app/com/wpbs/DTO"
)

type UserService struct {
	GetUser     func(ctx context.Context, req []interface{}, rsp *DTO.User) error
	CreateUser  func(ctx context.Context, req []interface{}, rsp *DTO.User) error
	GetUserById func(ctx context.Context, req []interface{}, rsp *DTO.User) error
}

func (u *UserService) Reference() string {
	return "UserService"
}
