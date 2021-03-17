package dao

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"user-service/app/com/wpbs/DTO"
)

var (
	insertMonoFileMock *DTO.MonoFile
	fileDao            *FileDao
)

func init() {
	fileDao = new(FileDao)

	insertMonoFileMock = &DTO.MonoFile{
		FileName:    "user.jpg",
		FileViewUrl: "ddddd",
		FileType:    "jpg",
		FileSize:    666,
		UserId:      160,
		UserName:    "安康",
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	}
}

func TestFileDao_CreateFile(t *testing.T) {
	err := fileDao.CreateFile(insertMonoFileMock)
	assert.Nil(t, err)
}
