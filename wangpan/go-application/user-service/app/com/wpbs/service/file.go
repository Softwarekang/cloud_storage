package service

import (
	"context"
	gxlog "github.com/dubbogo/gost/log"
	"user-service/app/com/wpbs/DTO"
	"user-service/app/com/wpbs/store"
)

// rpc service fileService
type FileService struct {
}

// uploadFile
func (f *FileService) UploadFile(ctx context.Context, req interface{}) (*DTO.MonoFile, error) {
	log.Infof("FileService UploadFile req:%v", req)
	session, err := store.DBClient.BeginTx()
	defer store.DBClient.EndTx(session, &err)

	monoFile := req.(*DTO.MonoFile)
	err = store.DBClient.File(session).CreateFile(monoFile)
	if err != nil {
		return nil, err
	}

	memory, err := store.DBClient.Memory(session).GetMemoryByUserId(monoFile.UserId)
	if err != nil {
		return nil, err
	}

	err = store.DBClient.Memory(session).UpdateMemory(monoFile.UserId, memory.ConsumeMemory+monoFile.FileSize)
	if err != nil {
		return nil, err
	}

	log.Infof("FileService UploadFile rsp:%v", *monoFile)
	return monoFile, nil
}

// getFile by userName
func (f *FileService) GetFileListByUserId(ctx context.Context, req interface{}) (*DTO.FileList, error) {
	log.Infof("GetFileListByUserId req:%v", req)
	getFileListReq := req.(*DTO.GetFileList)
	// 开启事务
	session, err := store.DBClient.BeginTx()
	if err != nil {
		return nil, err
	}
	defer store.DBClient.EndTx(session, &err)

	monoFiles, err := store.DBClient.File(session).GetFileListByUserID(getFileListReq.UserId, getFileListReq.FileType, getFileListReq.Page, getFileListReq.PageSize)
	if err != nil {
		return nil, err
	}
	totalCount, err := store.DBClient.File(session).GetTotalCountByUserId(getFileListReq.UserId, getFileListReq.FileType)
	if err != nil {
		return nil, err
	}
	rsp := &DTO.FileList{totalCount, monoFiles}
	log.Infof("FileService GetFileListByUserId rsp:%v", rsp)
	return rsp, nil
}

// delete files by ids
func (f *FileService) DeleteFileByIDs(ctx context.Context, req interface{}) error {
	gxlog.CInfo("FileService DeleteFileByIDs req:%v", req)
	engine := store.DBClient.Begin()
	deleteFile := req.(*DTO.DeleteFile)
	if err := store.DBClient.File(engine).DeleteFilesByIds(deleteFile.FileIds); err != nil {
		return err
	}
	return nil
}

// reference
func (f *FileService) Reference() string {
	return "FileService"
}
