package service

import (
	"errors"
	"gin-element-admin/global"
	"gin-element-admin/model"
	"gin-element-admin/model/request"
	"gorm.io/gorm"
)

// CreateApi 新增基础api
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error
func CreateApi(api model.SysApi) (err error) {
	if !errors.Is(global.GEA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GEA_DB.Create(&api).Error
}

// DeleteApi 删除基础api
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error
func DeleteApi(api model.SysApi) (err error) {
	err = global.GEA_DB.Delete(&api).Error
	ClearCasbin(1, api.Path, api.Method)
	return err
}

// GetAPIInfoList 分页获取数据
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: GetAPIInfoList
//@description: 分页获取数据
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: err error
func GetAPIInfoList(api model.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GEA_DB.Model(&model.SysApi{})
	var apiList []model.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

// GetAllApis 获取所有的api
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi
func GetAllApis() (err error, apis []model.SysApi) {
	err = global.GEA_DB.Find(&apis).Error
	return
}

// GetApiById 根据id获取api
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi
func GetApiById(id float64) (err error, api model.SysApi) {
	err = global.GEA_DB.Where("id = ?", id).First(&api).Error
	return
}

// UpdateApi 根据id更新api
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error
func UpdateApi(api model.SysApi) (err error) {
	var oldA model.SysApi
	err = global.GEA_DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.GEA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.GEA_DB.Save(&api).Error
		}
	}
	return err
}

// DeleteApisByIds 删除选中API
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: DeleteApis
//@description: 删除选中API
//@param: apis []model.SysApi
//@return: err error
func DeleteApisByIds(ids request.IdsReq) (err error) {
	err = global.GEA_DB.Delete(&[]model.SysApi{}, "id in ?", ids.Ids).Error
	return err
}
