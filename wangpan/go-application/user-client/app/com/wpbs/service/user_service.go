package service

import "context"

import (
	"user-client/app/com/wpbs/DTO"
)

type UserProvider struct {
	GetUser func(ctx context.Context, req []interface{}, rsp *DTO.User) error
}

func (u *UserProvider) Reference() string {
	return "UserProvider"
}
