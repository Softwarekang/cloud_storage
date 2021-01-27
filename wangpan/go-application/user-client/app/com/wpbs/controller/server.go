package controller

import (
	"context"
	"github.com/apache/dubbo-go/config"
	"user-client/app/com/wpbs/DTO"
	"user-client/app/com/wpbs/service"
	"user-client/common"
)

import (
	"github.com/gin-gonic/gin"
)

var (
	// 日志
	log = common.GetLogger()
	// ServerCheckService
	serverCheckService *service.ServerCheckService
)

// 服务探测
func CheckServer(ctx *gin.Context) {
	log.Info(" checkServer request:#v", ctx.Params)
	serverCheck := &DTO.ServerCheck{}
	log.Info("serverCheckService req:#v", "test")
	serverCheckService = config.GetConsumerService("ServerCheckService").(*service.ServerCheckService)
	err := serverCheckService.Check(context.TODO(), []interface{}{"echo test"}, serverCheck)
	if err != nil {
		log.Errorf("serverCheckService rsp error :#v", err)
		Error(ctx, common.CODE_500)
		return
	}

	log.Info("serverCheckService rsp :#v", serverCheck)
	if serverCheck.Code == 200 {
		Success(ctx, "server success", gin.H{"info": serverCheck})
	}
}
