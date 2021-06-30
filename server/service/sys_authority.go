package service

import (
	"errors"
	"gin-element-admin/global"
	"gin-element-admin/model"
	"gorm.io/gorm"
)

// CreateAuthority 创建一个角色
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth model.SysAuthority
//@return: err error, authority model.SysAuthority
func CreateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	var authorityBox model.SysAuthority
	if !errors.Is(global.GEA_DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), auth
	}
	err = global.GEA_DB.Create(&auth).Error
	return err, auth
}

// DeleteAuthority 删除角色
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: DeleteAuthority
//@description: 删除角色
//@param: auth *model.SysAuthority
//@return: err error
func DeleteAuthority(auth *model.SysAuthority) (err error) {
	if !errors.Is(global.GEA_DB.Where("authority_id = ?", auth.AuthorityId).First(&model.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GEA_DB.Where("parent_id = ?", auth.AuthorityId).First(&model.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := global.GEA_DB.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if len(auth.SysBaseMenus) > 0 {
		err = global.GEA_DB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		//err = db.Association("SysBaseMenus").Delete(&auth)
	} else {
		err = db.Error
	}
	ClearCasbin(0, auth.AuthorityId)
	return err
}
