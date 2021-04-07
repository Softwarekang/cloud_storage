package service

import "context"

import (
	"user-client/app/com/wpbs/DTO"
)

type UserService struct {
	GetUser       func(ctx context.Context, req []interface{}, rsp *DTO.User) error
	CreateUser    func(ctx context.Context, req []interface{}, rsp *DTO.User) error
	GetUserById   func(ctx context.Context, req []interface{}, rsp *DTO.User) error
	GetUserByName func(ctx context.Context, name string, rsp *DTO.User) error
	UpdateUser    func(ctx context.Context, req interface{}) error
}

func (u *UserService) Reference() string {
	return "UserService"
}
