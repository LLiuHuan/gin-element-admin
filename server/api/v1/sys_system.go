package v1

import (
	"gin-element-admin/global"
	"gin-element-admin/model/response"
	"gin-element-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetServerInfo 获取服务器信息
// @Tags System
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/getServerInfo [post]
func GetServerInfo(c *gin.Context) {
	if server, err := service.GetServerInfo(); err != nil {
		global.GEA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
	}
}

// GetSystemConfig 获取配置文件内容
// @Tags System
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/getSystemConfig [post]
func GetSystemConfig(c *gin.Context) {
	if err, config := service.GetSystemConfig(); err != nil {
		global.GEA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysConfigResponse{Config: config}, "获取成功", c)
	}
}
