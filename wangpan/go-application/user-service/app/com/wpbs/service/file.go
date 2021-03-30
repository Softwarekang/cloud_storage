package service

import (
	"context"
	gxlog "github.com/dubbogo/gost/log"
	"time"
	"user-service/app/com/wpbs/DTO"
	"user-service/app/com/wpbs/dao"
)

var (
	fileDao = new(dao.FileDao)
)

// rpc service fileService
type FileService struct {
}

// uploadFile
func (f *FileService) UploadFile(ctx context.Context, req interface{}) (*DTO.MonoFile, error) {
	gxlog.CInfo("req:%v", req)
	monoFile := req.(*DTO.MonoFile)
	monoFile.CreateTime = time.Now().Unix()
	monoFile.UpdateTime = time.Now().Unix()
	err := fileDao.CreateFile(monoFile)
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
	monoFiles, err := fileDao.GetFileListByUserID(getFileListReq.UserId, getFileListReq.FileType, getFileListReq.Page, getFileListReq.PageSize)
	if err != nil {
		log.Errorf("GetFileListByUserID error:%v", err)
		return nil, err
	}
	totalCount, err := fileDao.GetTotalCountByUserId(getFileListReq.UserId, getFileListReq.FileType)
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
	ids := req.([]int64)
	if err := fileDao.DeleteFilesByIds(ids); err != nil {
		log.Errorf("DeleteFilesByIds error:", err)
		return err
	}
	return nil
}

// reference
func (f *FileService) Reference() string {
	return "FileService"
}
