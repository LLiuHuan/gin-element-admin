package v1

import (
	"gin-element-admin/global"
	"gin-element-admin/model"
	"gin-element-admin/model/request"
	"gin-element-admin/model/response"
	"gin-element-admin/service"
	"gin-element-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateApi 创建基础api
// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/createApi [post]
func CreateApi(c *gin.Context) {
	var api model.SysApi
	_ = c.ShouldBindJSON(&api)
	if errStr, err := utils.BaseValidator(&api, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	if err := service.CreateApi(api); err != nil {
		global.GEA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteApi 删除api
// @Tags SysApi
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApi [post]
func DeleteApi(c *gin.Context) {
	var api model.SysApi
	if errStr, err := utils.BaseValidator(&api, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	if err := service.DeleteApi(api); err != nil {
		global.GEA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetApiList 分页获取API列表
// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func GetApiList(c *gin.Context) {
	var pageInfo request.SearchApiParams
	if errStr, err := utils.BaseValidator(&pageInfo, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	if err, list, total := service.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
		global.GEA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetApiById 根据id获取api
// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func GetApiById(c *gin.Context) {
	var idInfo request.GetById
	if errStr, err := utils.BaseValidator(&idInfo, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	err, api := service.GetApiById(idInfo.ID)
	if err != nil {
		global.GEA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(response.SysAPIResponse{Api: api}, c)
	}
}

// UpdateApi 创建基础api
// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /api/updateApi [post]
func UpdateApi(c *gin.Context) {
	var api model.SysApi
	if errStr, err := utils.BaseValidator(&api, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	if err := service.UpdateApi(api); err != nil {
		global.GEA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// GetAllApis 获取所有的Api 不分页
// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func GetAllApis(c *gin.Context) {
	if err, apis := service.GetAllApis(); err != nil {
		global.GEA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysAPIListResponse{Apis: apis}, "获取成功", c)
	}
}

// DeleteApisByIds 删除选中Api
// @Tags SysApi
// @Summary 删除选中Api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApisByIds [delete]
func DeleteApisByIds(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := service.DeleteApisByIds(ids); err != nil {
		global.GEA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
