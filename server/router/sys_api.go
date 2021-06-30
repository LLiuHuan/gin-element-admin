package router

import (
	v1 "gin-element-admin/api/v1"
	"gin-element-admin/middlewares"

	"github.com/gin-gonic/gin"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api").Use(middlewares.OperationRecord())
	{
		ApiRouter.POST("createApi", v1.CreateApi)               // 创建Api
		ApiRouter.POST("deleteApi", v1.DeleteApi)               // 删除Api
		ApiRouter.GET("getApiList", v1.GetApiList)              // 获取Api列表
		ApiRouter.GET("getApiById", v1.GetApiById)              // 获取单条Api消息
		ApiRouter.POST("updateApi", v1.UpdateApi)               // 更新api
		ApiRouter.GET("getAllApis", v1.GetAllApis)              // 获取所有api
		ApiRouter.DELETE("deleteApisByIds", v1.DeleteApisByIds) // 删除选中api
	}
}
