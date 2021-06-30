package service

import (
	"gin-element-admin/global"
	"gin-element-admin/model"
	"gin-element-admin/model/request"
)

// CreateSysOperationRecord 创建接口访问记录
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: CreateSysOperationRecord
//@description: 创建记录
//@param: sysOperationRecord model.SysOperationRecord
//@return: err error
func CreateSysOperationRecord(sysOperationRecord model.SysOperationRecord) (err error) {
	err = global.GEA_DB.Create(&sysOperationRecord).Error
	return err
}

// DeleteSysOperationRecordByIds 批量删除记录
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: DeleteSysOperationRecordByIds
//@description: 批量删除记录
//@param: ids request.IdsReq
//@return: err error
func DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	err = global.GEA_DB.Delete(&[]model.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

// DeleteSysOperationRecord 删除操作记录
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: DeleteSysOperationRecord
//@description: 删除操作记录
//@param: sysOperationRecord model.SysOperationRecord
//@return: err error
func DeleteSysOperationRecord(sysOperationRecord model.SysOperationRecord) (err error) {
	err = global.GEA_DB.Delete(&sysOperationRecord).Error
	return err
}

// GetSysOperationRecord 根据id获取单条操作记录
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: DeleteSysOperationRecord
//@description: 根据id获取单条操作记录
//@param: id uint
//@return: err error, sysOperationRecord model.SysOperationRecord
func GetSysOperationRecord(id uint) (err error, sysOperationRecord model.SysOperationRecord) {
	err = global.GEA_DB.Where("id = ?", id).First(&sysOperationRecord).Error
	return
}

// GetSysOperationRecordInfoList 分页获取操作记录列表
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: GetSysOperationRecordInfoList
//@description: 分页获取操作记录列表
//@param: info request.SysOperationRecordSearch
//@return: err error, list interface{}, total int64
func GetSysOperationRecordInfoList(info request.SysOperationRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GEA_DB.Model(&model.SysOperationRecord{})
	var sysOperationRecords []model.SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return err, sysOperationRecords, total
}
