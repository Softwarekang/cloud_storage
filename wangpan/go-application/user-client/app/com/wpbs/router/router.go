package router

import (
	"user-client/app/com/wpbs/controller"
	"user-client/common/cors"
	log2 "user-client/common/log"
)

import (
	"github.com/gin-gonic/gin"
)

var (
	// 日志
	log = log2.GetLogger()
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
	// upload image
	router.POST("/userImage", controller.UserImage)
	// 登录
	router.POST("/login", controller.Login)
	/*登录验证逻辑*/
	// get user
	router.GET("/user", controller.GetUserById)
	// 文件操作
	router.POST("/file", controller.FileUpload)
	router.POST("/getFile", controller.GetFileList)
	router.DELETE("/file", controller.DeleteFiles)
	// 内存操作
	router.GET("/memory", controller.GetMemory)
	// 视频
	router.GET("/video/:fileName", controller.GetVideo)
}
