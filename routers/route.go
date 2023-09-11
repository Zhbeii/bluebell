package routers

import (
	"github.com/gin-gonic/gin"
	"learn/bluebell/controllers"
	"learn/bluebell/logger"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	// 注册路由业务
	// todo 这个 signup 使用 postman 无法返回 200
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
