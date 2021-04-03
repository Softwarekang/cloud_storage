package minio

import (
	"errors"
	"github.com/minio/minio-go/v6"
	"io"
	"user-client/common/app_config"
	"user-client/common/log"
)

var (
	logger      = log.GetLogger()
	minioConfig *app_config.MinioConfig
)

// minio客户端封装
type Client struct {
	// config
	minioConfig app_config.MinioConfig
	// minio client
	minioClient *minio.Client
}

// init minio config
func init() {
	app := app_config.GetAppConfig()
	if app == nil {
		panic(errors.New("appConfig not init"))
	}
	// todo add default config
	minioConfig = &app.Minio
}

// 新连接
func (c *Client) NewMinIoClient() error {
	c.minioConfig = *minioConfig
	client, err := minio.NewWithRegion(c.minioConfig.EndPoint, c.minioConfig.AccessKeyID, c.minioConfig.SecretAccessKey,
		c.minioConfig.SSL, c.minioConfig.Region)
	if err != nil {
		logger.Error("connect minioClient fail:", err)
		return err
	}

	c.minioClient = client
	log.Info("minio client connect success:", c.minioConfig)
	return nil
}

// create bucket
func (c *Client) CreateBucket(bucketName string) error {
	check(c)
	err := c.minioClient.MakeBucket(bucketName, c.minioConfig.Region)
	if err != nil {
		logger.Error("minio client make bucket error msg:", err)
		return err
	}

	exists, err := c.minioClient.BucketExists(bucketName)
	if err != nil {
		logger.Error("minio client action error :", err)
		return err
	}

	if !exists {
		logger.Error("minio exists bucket not exists")
		return errors.New("make bucket fail")
	}

	return nil
}

// 检测桶是否存在
func (c *Client) ExistsBucket(bucketName string) (bool, error) {
	check(c)
	bucketExists, err := c.minioClient.BucketExists(bucketName)
	return bucketExists, err
}

// 构建viewUrl

// 创建文件
func (c *Client) UploadFile(bucketName, objectName, suffix string, file io.Reader, size int64) error {
	check(c)
	options := c.GetUploadOptions(suffix)
	length, err := c.minioClient.PutObject(bucketName, objectName, file, size, minio.PutObjectOptions{ContentType: options, NumThreads: 100})
	if err != nil {
		logger.Error("minio client put object error:", err)
		return err
	}

	logger.Info("minio client put byte size:", length)
	log.Infof("put fileInfo bucketName:%v, objectName:%v, suffix:%v,options:%v", bucketName, objectName, suffix, options)
	return nil
}

// 保证client init
func check(c *Client) {
	if &c.minioConfig == nil || c.minioClient == nil {
		c.NewMinIoClient()
	}
}

// 获取文件上传类型
func (c *Client) GetUploadOptions(suffix string) string {
	typeMap := make(map[string]string)
	/*// 特殊文本类型
	typeMap["html"] = "text/html"
	typeMap["xml"] = "text/xml"
	typeMap["json"] = "application/json"
	typeMap["pdf"] = "application/pdf"
	// 图片类型
	typeMap["gif"] = "image/gif"
	typeMap["jpeg"] = "image/jpeg"
	typeMap["png"] = "image/png"
	typeMap["jpg"] = "image/jpg"
	*/
	// 默认类型文件流形式
	typeMap["default"] = "application/octet-stream"

	if res, ok := typeMap[suffix]; ok {
		return res
	}

	return typeMap["default"]
}

/*
text/html ： HTML格式
text/plain ：纯文本格式
text/xml ： XML格式
image/gif ：gif图片格式
image/jpeg ：jpg图片格式
image/png：png图片格式
以application开头的媒体格式类型：

application/xhtml+xml ：XHTML格式
application/xml： XML数据格式
application/atom+xml ：Atom XML聚合格式
application/json： JSON数据格式
application/pdf：pdf格式
application/msword ： Word文档格式
application/octet-stream ： 二进制流数据（如常见的文件下载）
application/x-www-form-urlencoded ： <form encType=””>中默认的encType，form表单数据被编码为key/value格式发送到服务器（表单默认的提交数据的格式）
*/
