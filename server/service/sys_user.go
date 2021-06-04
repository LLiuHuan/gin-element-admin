package service

import (
	"gin-element-admin/global"
	"gin-element-admin/model"
	"gin-element-admin/utils"
)

// Login 用户登录
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser
func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GEA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}
