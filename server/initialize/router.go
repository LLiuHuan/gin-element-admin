package initialize

import (
	"gin-element-admin/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	var r = gin.Default()
	//r.Use(GinLogger(), GinRecovery(true))
	// 跨域
	r.Use(middlewares.Cors())

	// swagger 文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	PublicGroup := r.Group("")
	{
		PublicGroup.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
