package router

import (
	v1 "gin-element-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouter := Router.Group("system")
	{
		SystemRouter.POST("getServerInfo", v1.GetServerInfo)     // 获取服务器信息
		SystemRouter.POST("getSystemConfig", v1.GetSystemConfig) // 获取配置文件内容
	}
}
