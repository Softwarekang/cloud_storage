package common

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"sync"
	"user-client/common/app_config"
)

var (
	once   sync.Once
	Logger *logrus.Logger
	config = app_config.GetAppConfig()
)

// 获取 logger
func GetLogger() *logrus.Logger {
	once.Do(func() {
		loggerToFile()
	})

	return Logger
}

// 日志记录到文件
func loggerToFile() {
	logFilePath := config.Log.LogPath
	logFileName := config.Log.FileName
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	_, err := os.Stat(fileName)
	if (err != nil) {
		os.Create(fileName)
	}

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	Logger = logrus.New()

	// 设置输出
	Logger.SetOutput(src)

	// 设置日志级别
	Logger.SetLevel(logrus.DebugLevel)
}
