package controller

import (
	"context"
	"errors"
	"github.com/apache/dubbo-go/config"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strconv"
	"user-client/app/com/wpbs/BO"
	"user-client/app/com/wpbs/DTO"
	"user-client/app/com/wpbs/service"
	"user-client/common"
	"user-client/common/app_config"
	"user-client/common/httpcode"
	"user-client/common/minio"
	"user-client/common/utils"
)

// 文件操作controller
func FileUpload(ctx *gin.Context) {
	multiFile, userId, userName, err := checkFileUploadParam(ctx)
	if err != nil {
		logger.Error("fileUpload param error:", err)
		Error(ctx, httpcode.CODE_405)
		return
	}

	fileViewUrl, err := UploadFileMinio(multiFile)
	if err != nil {
		logger.Error("fileUpload controller uploadFileMinio error:", err)
		Error(ctx, httpcode.CODE_500)
		return
	}

	monoFile := &DTO.MonoFile{
		FileName:    multiFile.Filename,
		FileViewUrl: fileViewUrl,
		FileType:    utils.GetTypeBySuffix(utils.GetSuffix(multiFile.Filename)),
		FileSize:    multiFile.Size,
		UserId:      userId,
		UserName:    userName,
	}
	fileService := config.GetConsumerService("FileService").(*service.FileService)
	resFile := &DTO.MonoFile{}
	if err := fileService.UploadFile(context.TODO(), monoFile, resFile); err != nil {
		logger.Error("rpc call fileService error:", err)
		Error(ctx, httpcode.CODE_500)
		return
	}
	logger.Info("uploadFile success ")
	Success(ctx, "上传成功", gin.H{"info": resFile})
}

// 单文件上传参数校验
func checkFileUploadParam(ctx *gin.Context) (*multipart.FileHeader, int64, string, error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		logger.Errorf("file param err:", err)
		return nil, 0, "", err
	}

	Id := ctx.PostForm("userId")
	if Id == "" {
		return nil, 0, "", errors.New("param error userid is empty string")
	}

	userName := ctx.PostForm("userName")
	if userName == "" {
		return nil, 0, "", errors.New("param error userName is empty string")
	}
	userId, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		logger.Error(err)
		return nil, 0, "", err
	}
	return file, userId, userName, nil
}

// 文件上传封装
func UploadFileMinio(multiFile *multipart.FileHeader) (string, error) {
	minioClient := new(minio.Client)
	bucketName := common.BucketName
	existsBucket, _ := minioClient.ExistsBucket(bucketName)
	if !existsBucket {
		err := minioClient.CreateBucket(bucketName)
		if err != nil {
			logger.Error(err)
			return "", err
		}
	}

	fileName := multiFile.Filename
	fileSize := multiFile.Size
	suffix := utils.GetSuffix(fileName)
	minioServerUrl := app_config.GetAppConfig().Minio.EndPoint
	fileViewUrl := minioServerUrl + "/" + bucketName + "/" + fileName

	file, err := multiFile.Open()
	defer file.Close()
	if err != nil {
		logger.Error("file open error:", err)
		return "", err
	}

	err = minioClient.UploadFile(bucketName, fileName, suffix, file, fileSize)
	if err != nil {
		logger.Error("minio client uploadfile error:", err)
		return "", err
	}
	return fileViewUrl, nil
}

// 删除文件
func DeleteFiles(ctx *gin.Context) {
	array := ctx.QueryArray("ids")
	var ids []int64
	if err := utils.StringArrayToInt64Array(array, &ids); err != nil {
		logger.Errorf("DeleteFiles binding param error:%v", err)
		Error(ctx, httpcode.CODE_405)
		return
	}

	fileService := config.GetConsumerService("FileService").(*service.FileService)
	if err := fileService.DeleteFileByIDs(context.TODO(), ids); err != nil {
		logger.Errorf("rpc call fileService methods DeleteFileByIDs:%v", err)
		Error(ctx, httpcode.CODE_500)
		return
	}

	Success(ctx, "删除成功", nil)
}

// 获取文件列表
func GetFileList(ctx *gin.Context) {
	logger.Info(" controller GetFileList ")
	var getFileBo BO.GetFileList
	if err := ctx.ShouldBindJSON(&getFileBo); err != nil {
		logger.Errorf("getFileList bind param error:", err)
		Error(ctx, httpcode.CODE_405)
		return
	}

	fileService := config.GetConsumerService("FileService").(*service.FileService)
	req := &DTO.GetFileList{
		UserId:   getFileBo.UserId,
		FileType: getFileBo.FileType,
		Page:     getFileBo.Page,
		PageSize: getFileBo.PageSize,
	}
	fileList := &DTO.FileList{}
	if err := fileService.GetFileListByUserId(context.TODO(), req, fileList); err != nil {
		logger.Errorf("rpc fileService error:", err)
		Error(ctx, httpcode.CODE_500)
		return
	}

	Success(ctx, "获取成功", gin.H{"info": fileList})
}
