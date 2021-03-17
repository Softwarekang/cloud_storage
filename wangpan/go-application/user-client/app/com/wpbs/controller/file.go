package controller

import (
	"github.com/gin-gonic/gin"
)

// 文件操作controller
func FileUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		logger.Errorf("file param err:", err)
	}

	open, err := file.Open()
	defer open.Close()

}
