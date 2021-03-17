package dao

import (
	"user-service/app/com/wpbs/DTO"
	"user-service/app/com/wpbs/PO"
)

// file orm obj
type FileDao struct {
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
