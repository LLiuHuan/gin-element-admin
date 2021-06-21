package router

import (
	v1 "gin-element-admin/api/v1"
	"gin-element-admin/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middlewares.OperationRecord())
	{
		UserRouter.POST("register", v1.Register)             // 注册
		UserRouter.POST("changePassword", v1.ChangePassword) // 修改密码
		UserRouter.DELETE("delete", v1.DeleteUser)           // 删除用户
		UserRouter.POST("setAuthority", v1.SetAuthority)     // 设置用户权限
		UserRouter.PUT("setInfo", v1.SetUserInfo)            // 设置用户信息
		UserRouter.GET("list", v1.GetUserList)               // 分页获取用户列表
	}
}
