package dao

import (
	"time"
	"user-service/app/com/wpbs/DTO"
	"user-service/app/com/wpbs/PO"
	"user-service/common/extension"
)

// file orm obj
type FileDao struct {
	DB        interface{}
	SQLClient string
}

func NewFileDao(DB interface{}, arg ...string) *FileDao {
	fileDao := &FileDao{}
	if arg != nil {
		fileDao.SQLClient = arg[0]
	}
	fileDao.DB = DB
	return fileDao
}

// 插入文件
func (fileDao *FileDao) CreateFile(monoFile *DTO.MonoFile) error {
	log.Infof("FileDao CreateFile param:%v", *monoFile)
	monoFile.CreateTime = time.Now().Unix()
	monoFile.UpdateTime = time.Now().Unix()
	fileModel := ChangeMonoFileVP(monoFile)
	client := extension.GetSQLClient(fileDao.SQLClient)
	if _, err := client.Insert(fileDao.DB, fileModel); err != nil {
		log.Error("FileDao CreateFile error:%v, model:%v", err, *fileModel)
		return err
	}

	log.Info("FileDao CreateFile success rsp model:%v", *fileModel)
	return nil
}

// 根据用户ID查询文件列表
func (fileDao *FileDao) GetFileListByUserID(userId int64, fileType string, page, pageSize int) ([]*DTO.MonoFile, error) {
	sql := WithSQLParam("select * from file", "user_id", userId)
	sql = WithSQLParam(sql, "file_type", fileType)
	var monoFileList []*PO.File
	offset := (page - 1) * pageSize
	client := extension.GetSQLClient(fileDao.SQLClient)
	if err := client.SQL(fileDao.DB, sql).Limit(pageSize, offset).Find(&monoFileList); err != nil {
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
	client := extension.GetSQLClient(fileDao.SQLClient)
	var totalCount int
	_, err := client.SQL(fileDao.DB, sql).Get(&totalCount)
	if err != nil {
		log.Errorf("GetTotalCountByUserId  error:%v", err)
		return 0, nil
	}

	return totalCount, nil
}

// 删除file
func (fileDao *FileDao) DeleteFilesByIds(ids []int64) error {
	sql := "delete from file where id in " + IDToStr(ids)
	client := extension.GetSQLClient(fileDao.SQLClient)
	_, err := client.SQL(fileDao.DB, sql).Execute()
	if err != nil {
		log.Errorf("DeleteFilesByIds error:%v", err)
		return err
	}

	return nil
}

func (fileDao *FileDao) GetFileListByIds(id []int64) ([]*DTO.MonoFile, error) {
	log.Infof("FileDao GetFileListByIds param:Ids:%v", id)
	sql := "select * from file where id in " + IDToStr(id)
	client := extension.GetSQLClient(fileDao.SQLClient)
	var monoFileList []*PO.File
	if err := client.SQL(fileDao.DB, sql).Find(&monoFileList); err != nil {
		log.Errorf("FileDao GetFileListByIds error:%v, sql:%v", err, sql)
		return nil, err
	}

	resFileList := make([]*DTO.MonoFile, 0, len(monoFileList))
	for _, file := range monoFileList {
		resFileList = append(resFileList, ChangeMonoFilePV(file))
	}
	log.Infof("FileDao GetFileListByIds  success rsp model:%v", resFileList)
	return resFileList, nil
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
