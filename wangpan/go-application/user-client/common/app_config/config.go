package app_config

import (
	"gopkg.in/ini.v1"
	"os"
	"sync"
)

var (
	config *configStruct
	once   sync.Once
	App    appStruct

	// 环境配置文件
	ENV_CONFIG_FILE = os.Getenv("ENV_CONFIG_FILE")
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

// 加载环境配置信息
func LoadConfig() {
	cfg, err := ini.LooseLoad(ENV_CONFIG_FILE)
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
		LoadConfig()
	})

	return config
}
