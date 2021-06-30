package v1

import (
	"gin-element-admin/global"
	"gin-element-admin/model/request"
	"gin-element-admin/model/response"
	"gin-element-admin/service"
	"gin-element-admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateCasbin 更新角色api权限
// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /casbin/UpdateCasbin [post]
func UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInReceive
	if errStr, err := utils.BaseValidator(&cmr, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	if err := service.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos); err != nil {
		global.GEA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetPolicyPathByAuthorityId 获取权限列表
// @Tags Casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func GetPolicyPathByAuthorityId(c *gin.Context) {
	var casbin request.CasbinInReceive
	if errStr, err := utils.BaseValidator(&casbin, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	paths := service.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	response.OkWithDetailed(response.PolicyPathResponse{Paths: paths}, "获取成功", c)
}
