package router

import (
	v1 "gin-element-admin/api/v1"
	"gin-element-admin/middlewares"

	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	CasbinRouter := Router.Group("casbin").Use(middlewares.OperationRecord())
	{
		CasbinRouter.POST("updateCasbin", v1.UpdateCasbin)                             // 更新权限
		CasbinRouter.POST("getPolicyPathByAuthorityId", v1.GetPolicyPathByAuthorityId) // 获取权限列表
	}
}
