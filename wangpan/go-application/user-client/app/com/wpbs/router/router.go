package router

import (
	"user-client/app/com/wpbs/controller"
	"user-client/common"
)

import (
	"github.com/gin-gonic/gin"
)

var (
	// 日志
	log = common.Logger
)

// 加载router
func LoadRouters(router *gin.Engine) {
	log.Info("router loading ")
	gin.DisableConsoleColor()
	//404错误
	router.NoRoute(controller.NoRoute)
	//server check
	router.GET("/test", controller.CheckServer)
	log.Info("router load success")
}
