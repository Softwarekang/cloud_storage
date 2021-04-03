package controller

import (
	"context"
	"github.com/apache/dubbo-go/config"
	"github.com/gin-gonic/gin"
	"strconv"
	"user-client/app/com/wpbs/DTO"
	"user-client/app/com/wpbs/service"
	"user-client/common/httpcode"
)

func GetMemory(ctx *gin.Context) {
	logger.Info("controller GetMemory handing")
	var userId int64
	if err := checkGetMemoryParam(ctx.Query("userId"), &userId); err != nil {
		logger.Errorf("controller GetMemory binding param error:%v", err)
		Error(ctx, httpcode.CODE_405)
		return
	}

	memoryService := config.GetConsumerService("MemoryService").(*service.MemoryService)
	var memory DTO.Memory
	if err := memoryService.GetMemory(context.TODO(), userId, &memory); err != nil {
		logger.Errorf("rpc call memoryService error:%v", err)
		Error(ctx, httpcode.CODE_500)
		return
	}

	logger.Infof("controller GetMemory success rsp model:%v", memory)
	Success(ctx, "获取成功", gin.H{"info": memory})
}

func checkGetMemoryParam(param string, userId *int64) error {
	parseInt, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		logger.Errorf("GetMemoryParam userId must be number type int")
		return err
	}

	*userId = parseInt
	return nil
}
