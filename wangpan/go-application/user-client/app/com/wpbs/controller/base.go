package controller

import (
	"net/http"
	"user-client/common/httpcode"
)

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Code    httpcode.ResponseCode `json:"code"`
	Message string                `json:"message"`
	Data    interface{}           `json:"data"`
}

func Success(ctx *gin.Context, message string, data map[string]interface{}) {

	response := response{
		Code:    httpcode.SUCCESS,
		Message: message,
		Data:    data,
	}
	setResponse(ctx, http.StatusOK, response)
}

func Error(ctx *gin.Context, code httpcode.ResponseCode) {

	response := response{
		Code:    code,
		Message: httpcode.GetCodeText(code),
		Data:    gin.H{},
	}
	setResponse(ctx, http.StatusOK, response)
}

func ErrorWithMessage(ctx *gin.Context, code httpcode.ResponseCode, message string) {

	response := response{
		Code:    code,
		Message: message,
		Data:    gin.H{},
	}
	setResponse(ctx, http.StatusOK, response)
}

func setResponse(ctx *gin.Context, statusCode int, resp response) {
	ctx.Set("response", resp)
	ctx.JSON(statusCode, resp)
}

//NOTFOUND method not found action
func NOTFOUND(ctx *gin.Context) {
	response := response{
		Code:    httpcode.CODE_404,
		Message: httpcode.GetCodeText(httpcode.CODE_404),
		Data:    gin.H{},
	}
	ctx.Set("response", response)
	ctx.JSON(http.StatusNotFound, response)
}

//StatusInternalServerError server 500 error
func StatusInternalServerError(ctx *gin.Context) {
	response := response{
		Code:    httpcode.CODE_500,
		Message: httpcode.GetCodeText(httpcode.CODE_500),
		Data:    gin.H{},
	}
	ctx.Set("response", response)
	ctx.JSON(http.StatusInternalServerError, response)
}
