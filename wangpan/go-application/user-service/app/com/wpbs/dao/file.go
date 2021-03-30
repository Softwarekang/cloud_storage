package dao

import (
	"user-service/app/com/wpbs/DTO"
	"user-service/app/com/wpbs/PO"
)

// file orm obj
type FileDao struct {
	DB interface{}
}

func NewFileDao(DB interface{}) *FileDao {
	fileDao := new(FileDao)
	fileDao.DB = DB
	return fileDao
}
// 插入文件
func (fileDao *FileDao) CreateFile(monoFile *DTO.MonoFile) error {
	fileModel := ChangeMonoFileVP(monoFile)
	if _, err := DB.Insert(fileModel); err != nil {
		log.Error("monofile insert error msg:", err)
		return err
	}

	log.Info("monoFile insert success msg:", fileModel)
	return nil
}

// 根据用户ID查询文件列表
func (fileDao *FileDao) GetFileListByUserID(userId int64, fileType string, page, pageSize int) ([]*DTO.MonoFile, error) {
	sql := WithSQLParam("select * from file", "user_id", userId)
	sql = WithSQLParam(sql, "file_type", fileType)
	var monoFileList []*PO.File
	offset := (page - 1) * pageSize
	if err := DB.SQL(sql).Limit(pageSize, offset).Find(&monoFileList); err != nil {
		log.Errorf("file dao GetFileListByUserID sql :%v error:%v", sql, err)
		return nil, err
	}

	resFileList := make([]*DTO.MonoFile, 0, len(monoFileList))
	for _, file := range monoFileList {
		resFileList = append(resFileList, ChangeMonoFilePV(file))
	}

	return resFileList, nil
}

// 查询total count
func (fileDao *FileDao) GetTotalCountByUserId(userId int64, fileType string) (int, error) {
	sql := WithSQLParam("select count(id) from file", "user_id", userId)
	sql = WithSQLParam(sql, "file_type", fileType)
	var totalCount int
	_, err := DB.SQL(sql).Get(&totalCount)
	if err != nil {
		log.Errorf("GetTotalCountByUserId  error:%v", err)
		return 0, nil
	}

	return totalCount, nil
}

// 删除file
func (fileDao *FileDao) DeleteFilesByIds(ids []int64) error {
	sql := "delete from file where id in " + IDToStr(ids)
	_, err := DB.SQL(sql).Execute()
	if err != nil {
		log.Errorf("DeleteFilesByIds error:%v", err)
		return err
	}

	return nil
}
func ChangeMonoFileVP(monoFile *DTO.MonoFile) *PO.File {
	return &PO.File{
		Id:          monoFile.Id,
		FileName:    monoFile.FileName,
		FileViewUrl: monoFile.FileViewUrl,
		FileType:    monoFile.FileType,
		FileSize:    monoFile.FileSize,
		UserId:      monoFile.UserId,
		UserName:    monoFile.UserName,
		CreateTime:  monoFile.CreateTime,
		UpdateTime:  monoFile.UpdateTime,
	}
}

func ChangeMonoFilePV(monoFile *PO.File) *DTO.MonoFile {
	return &DTO.MonoFile{
		Id:          monoFile.Id,
		FileName:    monoFile.FileName,
		FileViewUrl: monoFile.FileViewUrl,
		FileType:    monoFile.FileType,
		FileSize:    monoFile.FileSize,
		UserId:      monoFile.UserId,
		UserName:    monoFile.UserName,
		CreateTime:  monoFile.CreateTime,
		UpdateTime:  monoFile.UpdateTime,
	}
}
