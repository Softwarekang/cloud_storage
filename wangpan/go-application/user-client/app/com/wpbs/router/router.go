package router

import (
	"user-client/app/com/wpbs/controller"
	"user-client/common"
	"user-client/common/cors"
)

import (
	"github.com/gin-gonic/gin"
)

var (
	// 日志
	log = common.GetLogger()
)

// 加载router
func LoadRouters(router *gin.Engine) {
	log.Info("router loading ")
	gin.DisableConsoleColor()
	/*跨域中间件*/
	router.Use(cors.Cors())
	//404错误
	router.NoRoute(controller.NoRoute)
	//server check
	router.GET("/test", controller.CheckServer)
	// create user
	router.POST("/user", controller.CreateUser)

	/*登录验证逻辑*/
	// get user
	router.GET("/user", controller.GetUserById)
	log.Info("router load success")
}
