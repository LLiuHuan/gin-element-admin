package router

import (
	v1 "gin-element-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)     // 登录
		BaseRouter.POST("captcha", v1.Captcha) // 获取验证码
		//BaseRouter.POST("test", v1.Test)
	}
	return BaseRouter
}
