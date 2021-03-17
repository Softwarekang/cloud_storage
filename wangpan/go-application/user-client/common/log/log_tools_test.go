package common

import (
	"gotest.tools/assert"
	"testing"
)

// 测试日志输出
func TestGetLogger(t *testing.T) {
	logger1 := GetLogger()
	logger2 := GetLogger()
	logger1.Infof("%v test log", 11)
	assert.Equal(t, logger1, logger2)
}