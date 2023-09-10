package routers

import (
	"github.com/gin-gonic/gin"
	"learn/bluebell/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册路由信息
	return r
}
