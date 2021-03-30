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

func TestFileDao_GetFileListByUserID(t *testing.T) {
	monoFiles, err := fileDao.GetFileListByUserID(160, "", 1, 8)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(monoFiles))
}

func TestFileDao_GetTotalCountByUserId(t *testing.T) {
	count, err := fileDao.GetTotalCountByUserId(160, "")
	assert.Nil(t, err)
	assert.Equal(t, 3, count)
}

func TestFileDao_DeleteFilesByIds(t *testing.T) {
	id := make([]int64, 0, 2)
	id = append(id, 1)
	id = append(id, 160)
	id = append(id, 4)
	id = append(id, 5)
	err := fileDao.DeleteFilesByIds(id)
	assert.Nil(t, err)
}
