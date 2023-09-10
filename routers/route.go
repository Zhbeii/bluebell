package routers

import (
	"github.com/gin-gonic/gin"
	"learn/bluebell/controllers"
	"learn/bluebell/logger"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册路由业务
	r.POST("/signup", controllers.SignUpHandler)
	r.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
