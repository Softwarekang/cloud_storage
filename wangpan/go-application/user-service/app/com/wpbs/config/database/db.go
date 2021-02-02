package database

import (
	"sync"
	"time"
	"user-service/common"
	"user-service/common/app_config"
)
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var (
	engine    *xorm.Engine
	once      sync.Once
	appConfig = app_config.GetAppConfig()
	log       = common.GetLogger()
)

// 获取DB
func GetDB() *xorm.Engine {
	once.Do(func() {
		initDB()
	})

	return engine
}

//  load db config
func initDB() {
	log.Info("init db connection")
	engineName := appConfig.DB.Engine
	dataBaseUrl := appConfig.DB.DATABASE_URL
	var err error
	engine, err = xorm.NewEngine(engineName, dataBaseUrl)
	if err != nil {
		log.Error("db connection  error driver", dataBaseUrl)
		panic(err)
	}


	// set logger
	engine.SetLogger(log)
	// set maxIdleConns
	engine.SetMaxIdleConns(app_config.DB.MaxIdleConns)
	// set MaxOpenConns
	engine.SetMaxOpenConns(app_config.DB.MaxOpenConns)
	// set ConnMaxLifeTime
	engine.SetConnMaxLifetime(time.Duration(app_config.DB.ConnMaxLifeTime))

	log.Info("database load success")

}
