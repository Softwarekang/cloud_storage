package app_config

import (
	"gotest.tools/assert"
	"testing"
)

// 全局配置加载测试
func TestConfigLoad(t *testing.T) {
	LoadConfig()
    // assert logPath
	assert.Equal(t, config.Log.LogPath, "D:/all-log/dubbo-go/dubbo-go")
	assert.Equal(t, config.App.AppName, "UserClient")
	assert.Equal(t, config.Minio.SSL, false)
}

// 单列测试
func TestGetAppConfig(t *testing.T) {
	appConfig1 := GetAppConfig()
	appConfig2 := GetAppConfig()
	assert.Equal(t, appConfig1, appConfig2)
}
