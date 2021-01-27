package app_config

import (
	"gopkg.in/ini.v1"
	"sync"
)

var (
	config *configStruct
	once   sync.Once
)

// 全部配置 app redis Log DB
type configStruct struct {
	App   appStruct   `ini:"app"`
	Redis redisStruct `ini:"redis"`
	Log   logStruct   `ini:"log"`
	DB    dbStruct    `ini:"database"`
}

//应用配置
type appStruct struct {
	AppName   string `ini:"app_name"`
	Port      string `ini:"port"`
	JWT_TOKEN string `ini:"jwt_token"`
}

var App appStruct

// 加载环境配置信息
func loadConfig() {
	cfg, err := ini.LooseLoad("D:\\GoFiles\\project-conf\\dev\\.env")
	if err != nil {
		panic(err)
	}

	config = new(configStruct)
	err = cfg.MapTo(config)
	if err != nil {
		panic(err)
	}

	App = config.App
	Redis = config.Redis
	Log = config.Log
	DB = config.DB
}

// 单列模式获取config
func GetAppConfig() *configStruct {
	once.Do(func() {
		loadConfig()
	})

	return config
}
