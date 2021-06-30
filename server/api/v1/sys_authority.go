package v1

import (
	"gin-element-admin/global"
	"gin-element-admin/model"
	"gin-element-admin/model/response"
	"gin-element-admin/service"
	"gin-element-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateAuthority
// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authority/createAuthority [post]
func CreateAuthority(c *gin.Context) {
	var authority model.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	if errStr, err := utils.BaseValidator(&authority, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	if err, authBack := service.CreateAuthority(authority); err != nil {
		global.GEA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
	}
}

// DeleteAuthority 删除角色
// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authority/deleteAuthority [post]
func DeleteAuthority(c *gin.Context) {
	var authority model.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	if errStr, err := utils.BaseValidator(&authority, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	if err := service.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.GEA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
