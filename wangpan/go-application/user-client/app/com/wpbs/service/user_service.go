package service

import "context"

import (
	"user-client/app/com/wpbs/pojo"
)

type UserProvider struct {
	GetUser func(ctx context.Context, req []interface{}, rsp *pojo.User) error
}

func (u *UserProvider) Reference() string {
	return "UserProvider"
}
