package service

import (
	"context"
	gxlog "github.com/dubbogo/gost/log"
	"time"
	"user-service/app/com/wpbs/DTO"
	"user-service/app/com/wpbs/store"
)

// rpc service fileService
type FileService struct {
}

// uploadFile
func (f *FileService) UploadFile(ctx context.Context, req interface{}) (*DTO.MonoFile, error) {
	gxlog.CInfo("req:%v", req)
	engine := store.DBClient.Begin()
	monoFile := req.(*DTO.MonoFile)
	monoFile.CreateTime = time.Now().Unix()
	monoFile.UpdateTime = time.Now().Unix()
	err := store.DBClient.File(engine).CreateFile(monoFile)
	if err != nil {
		log.Error("fileService uploadFile error:", err)
		return nil, err
	}

	log.Info("fileService uploadFile success")
	gxlog.CInfo("rsp:%v", monoFile)
	return monoFile, nil
}

// getFile by userName
func (f *FileService) GetFileListByUserId(ctx context.Context, req interface{}) (*DTO.FileList, error) {
	gxlog.CInfo("GetFileListByUserId req:%v", req)
	getFileListReq := req.(*DTO.GetFileList)
	// 开启事务
	var err error
	session, err := store.DBClient.BeginTx()
	if err != nil {
		return nil, err
	}
	defer store.DBClient.EndTx(session, &err)
	monoFiles, err := store.DBClient.File(session).GetFileListByUserID(getFileListReq.UserId, getFileListReq.FileType, getFileListReq.Page, getFileListReq.PageSize)
	if err != nil {
		log.Errorf("GetFileListByUserID error:%v", err)
		return nil, err
	}
	totalCount, err := store.DBClient.File(session).GetTotalCountByUserId(getFileListReq.UserId, getFileListReq.FileType)
	if err != nil {
		log.Errorf("GetTotalCountByUserId error:%v", err)
		return nil, err
	}
	rsp := &DTO.FileList{totalCount, monoFiles}
	return rsp, nil
}

// delete files by ids
func (f *FileService) DeleteFileByIDs(ctx context.Context, req interface{}) error {
	gxlog.CInfo("req:%v", req)
	engine := store.DBClient.Begin()
	deleteFile := req.(*DTO.DeleteFile)
	if err := store.DBClient.File(engine).DeleteFilesByIds(deleteFile.FileIds); err != nil {
		log.Errorf("DeleteFilesByIds error:%v", err)
		return err
	}
	return nil
}

// reference
func (f *FileService) Reference() string {
	return "FileService"
}
