package initialize

import (
	"gin-element-admin/global"
	"gin-element-admin/middlewares"
	"gin-element-admin/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

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

	// 令牌桶 限流
	r.Use(middlewares.RateLimitMiddleware(time.Second))
	// IP 限流
	if global.GEA_CONFIG.RateLimit.IpVerify {
		r.Use(middlewares.IpVerifyMiddleware())
	}

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
		router.InitSystemRouter(PrivateGroupV1)    // system相关路由
		router.InitUserRouter(PrivateGroupV1)      // 用户相关路由
		router.InitAuthorityRouter(PrivateGroupV1) // 角色相关路由
		router.InitApiRouter(PrivateGroupV1)       // Api相关路由
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
