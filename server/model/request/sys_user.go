package request

import uuid "github.com/satori/go.uuid"

// Login 用户登录结构体
type Login struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`   // 验证码
	CaptchaId string `json:"captchaId" binding:"required"` // 验证码ID
}

// Register 注册用户结构体
type Register struct {
	Username    string `json:"userName" binding:"required"`
	Password    string `json:"passWord" binding:"required"`
	NickName    string `json:"nickName" gorm:"default:'GEATest'" binding:"required"`
	HeaderImg   string `json:"headerImg" gorm:"default:''"`
	AuthorityId string `json:"authorityId" gorm:"default:888" binding:"required"`
}

// ChangePasswordStruct 修改密码结构体
type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// SetUserAuth Modify  user's auth structure
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`        // 用户UUID
	AuthorityId string    `json:"authorityId"` // 角色ID
}
