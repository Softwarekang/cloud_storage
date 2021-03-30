package service

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"user-service/app/com/wpbs/DTO"
)

var fileService = new(FileService)

func TestFileService_GetFileListByUserId(t *testing.T) {
	req := &DTO.GetFileList{
		UserId:   160,
		Page:     1,
		PageSize: 2,
	}

	rsp, err := fileService.GetFileListByUserId(context.TODO(), req)
	fmt.Print(*rsp.FileList[0])
	assert.Nil(t, err)
	assert.Equal(t, rsp.TotalCount, 3)
}

