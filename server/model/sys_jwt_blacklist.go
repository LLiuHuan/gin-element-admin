package model

import (
	"gin-element-admin/global"
)

type JwtBlacklist struct {
	global.GEA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
