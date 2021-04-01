package service

import (
	"context"
	"strconv"
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
	log.Infof("FileService DeleteFileByIDs req:%v", req)
	deleteFile := req.(*DTO.DeleteFile)
	session, err := store.DBClient.BeginTx()
	defer store.DBClient.EndTx(session, &err)

	fileList, err := store.DBClient.File(session).GetFileListByIds(deleteFile.FileIds)
	if err != nil {
		return err
	}

	err = store.DBClient.File(session).DeleteFilesByIds(deleteFile.FileIds)
	if err != nil {
		return err
	}

	totalFileSize := getTotalFileSize(fileList)
	userId, _ := strconv.ParseInt(deleteFile.UserId, 10, 64)
	memory, err := store.DBClient.Memory(session).GetMemoryByUserId(userId)
	if err != nil {
		return err
	}

	err = store.DBClient.Memory(session).UpdateMemory(userId, memory.ConsumeMemory-totalFileSize)
	if err != nil {
		return err
	}

	log.Infof("FileService DeleteFileByIDs success")
	return nil
}

func getTotalFileSize(fileList []*DTO.MonoFile) int64 {
	var total int64
	for _, v := range fileList {
		total += v.FileSize
	}

	return total
}

// reference
func (f *FileService) Reference() string {
	return "FileService"
}
