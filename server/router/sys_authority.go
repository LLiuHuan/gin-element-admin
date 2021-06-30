package router

import (
	v1 "gin-element-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAuthorityRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	AuthorityRouter := Router.Group("authority")
	{
		AuthorityRouter.POST("roles", v1.CreateAuthority)   // 创建角色
		AuthorityRouter.DELETE("roles", v1.DeleteAuthority) // 删除角色

	}
	return AuthorityRouter
}
