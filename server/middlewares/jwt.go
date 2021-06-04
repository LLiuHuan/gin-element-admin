package middlewares

import (
	"gin-element-admin/global"
	"gin-element-admin/model/request"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWT 结构体
type JWT struct {
	SigningKey []byte // 唯一签名
}

// NewJWT 创建JWT对象
func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GEA_CONFIG.JWT.SigningKey),
	}
}

// CreateJWT 创建token
func (j *JWT) CreateJWT(claims request.CustomClaims) (accessToken, refreshToken string, err error) {
	// access_token 短token
	// refresh_token 长token

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(global.GEA_CONFIG.JWT.RefreshExpiresTime)).Unix(),
	}).SignedString(j.SigningKey)
	if err != nil {
		return
	}
	claims.RefreshToken = refreshToken
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SigningKey)
	if err != nil {
		return
	}
	return
}
