package controller

import (
	"context"
	"github.com/apache/dubbo-go/config"
	"user-client/app/com/wpbs/DTO"
	"user-client/app/com/wpbs/service"
	"user-client/common/httpcode"
	"user-client/common/log"
)

import (
	"github.com/gin-gonic/gin"
)

var (
	// 日志
	logger = log.GetLogger()
	// ServerCheckService
	serverCheckService *service.ServerCheckService
)

// 服务探测
func CheckServer(ctx *gin.Context) {
	logger.Info(" checkServer request:", ctx.Params)
	serverCheck := &DTO.ServerCheck{}
	logger.Info("serverCheckService req:", "test")
	serverCheckService = config.GetConsumerService("ServerCheckService").(*service.ServerCheckService)
	err := serverCheckService.Check(context.TODO(), []interface{}{"echo test"}, serverCheck)
	if err != nil {
		logger.Errorf("serverCheckService rsp error :#v", err)
		Error(ctx, httpcode.CODE_500)
		return
	}

	logger.Info("serverCheckService rsp :#v", serverCheck)
	if serverCheck.Code == 200 {
		Success(ctx, "server success", gin.H{"info": serverCheck})
	}
}
