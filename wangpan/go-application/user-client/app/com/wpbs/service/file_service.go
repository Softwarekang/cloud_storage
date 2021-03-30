package service

import (
	"context"
	"user-client/app/com/wpbs/DTO"
)

// consumer rpc service
type FileService struct {
	// req:基本类型随便传, struct类型必须为 ptr   必须有rsp
	UploadFile          func(ctx context.Context, req interface{}, rsp *DTO.MonoFile) error
	GetFileListByUserId func(ctx context.Context, req interface{}, rsp *DTO.FileList) error
	DeleteFileByIDs     func(ctx context.Context, req interface{}) error
}

func (f *FileService) Reference() string {
	return "FileService"
}
