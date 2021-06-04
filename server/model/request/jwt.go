package request

import (
	"gin-element-admin/model"
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	RefreshToken string
	UserInfo     model.SysUser
	BufferTime   int64
	jwt.StandardClaims
}
