package router

import (
	v1 "gin-element-admin/api/v1"
	"gin-element-admin/middlewares"

	"github.com/gin-gonic/gin"
)

func InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	SysOperationRecordRouter := Router.Group("sysOperationRecord").Use(middlewares.OperationRecord())
	{
		SysOperationRecordRouter.POST("createSysOperationRecord", v1.CreateSysOperationRecord)             // 新建操作记录
		SysOperationRecordRouter.DELETE("deleteSysOperationRecord", v1.DeleteSysOperationRecord)           // 删除操作记录
		SysOperationRecordRouter.DELETE("deleteSysOperationRecordByIds", v1.DeleteSysOperationRecordByIds) // 批量删除操作记录
		SysOperationRecordRouter.GET("findSysOperationRecord", v1.FindSysOperationRecord)                  // 根据ID获取操作记录
		SysOperationRecordRouter.GET("getSysOperationRecordList", v1.GetSysOperationRecordList)            // 获取操作记录列表

	}
}
