package initialize

import (
	"gin-element-admin/global"
	"gin-element-admin/middlewares"
	"gin-element-admin/router"
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
	// Router.Use(middleware.LoadTls()  // https
	global.GEA_LOG.Info("use middleware logger")
	// 跨域
	r.Use(middlewares.Cors())

	global.GEA_LOG.Info("use middleware cors")
	// swagger 文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	global.GEA_LOG.Info("register swagger handler")

	PublicGroup := r.Group("/v1")
	{
		router.InitBaseRouter(PublicGroup)
	}

	PrivateGroupV1 := r.Group("/v1")
	PrivateGroupV1.Use(middlewares.JWTAuto()).Use(middlewares.CasbinHandler())
	{
		PrivateGroupV1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		router.InitSystemRouter(PrivateGroupV1)
		router.InitUserRouter(PrivateGroupV1)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
