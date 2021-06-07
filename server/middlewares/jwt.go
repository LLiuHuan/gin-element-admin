package middlewares

import (
	"errors"
	"fmt"
	"gin-element-admin/global"
	"gin-element-admin/model"
	"gin-element-admin/model/request"
	"gin-element-admin/model/response"
	"gin-element-admin/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"time"
)

var (
	TokenExpired     = errors.New("Token is expired")            // token 已过期
	TokenNotValidYet = errors.New("Token not active yet")        // token未激活
	TokenMalformed   = errors.New("That's not even a token")     // 不是token
	TokenInvalid     = errors.New("Couldn't handle this token:") // 无法处理token
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

// CreateToken 创建token
func (j *JWT) CreateToken(claims request.CustomClaims) (accessToken string, err error) {
	// access_token 短token
	// refresh_token 长token

	if claims.RefreshToken == "" {
		refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(global.GEA_CONFIG.JWT.RefreshExpiresTime)).Unix(),
		}).SignedString(j.SigningKey)
		if err != nil {
			return "", err
		}
		claims.RefreshToken = refreshToken
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SigningKey)
	if err != nil {
		return
	}
	return
}

// ParseToken 解析 token
func (j *JWT) ParseToken(accessToken string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(accessToken string) (newAccessToken string, err error) {
	// 获取accessToken的数据
	//claims, err := j.ParseToken(accessToken)
	var claims request.CustomClaims
	_, parseErr := jwt.ParseWithClaims(accessToken, &claims, func(*jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	// RefreshToken 过期
	if _, err = jwt.Parse(claims.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	}); err != nil {
		return
	}

	if v, _ := parseErr.(*jwt.ValidationError); v.Errors == jwt.ValidationErrorExpired {
		jwt.TimeFunc = func() time.Time {
			return time.Unix(0, 0)
		}
		// 刷新AccessToken
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(global.GEA_CONFIG.JWT.AccessExpiresTime)).Unix()
		return j.CreateToken(claims)
	}
	return accessToken, nil
}

func JWTAuto() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.FailWithMessage("请求头中auth格式有误", c)
			c.Abort()
			return
		}
		token := parts[1]
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if service.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			fmt.Println("token 授权已过期")
			if err == TokenExpired {
				accessToken, _ := j.RefreshToken(token)
				if accessToken == "" {
					response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
					c.Abort()
					return
				}
				c.Header("Authorization", "Bearer "+accessToken)
				newClaims, _ := j.ParseToken(accessToken)

				// TODO: 看看到时候把这里去掉，有点浪费时间
				if err, _ = service.FindUserByUuid(newClaims.UserInfo.UUID.String()); err != nil {
					_ = service.JsonInBlacklist(model.JwtBlacklist{Jwt: token})
					response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
					c.Abort()
				}

				if global.GEA_CONFIG.System.UseMultipoint {
					err, RedisJwtToken := service.GetRedisJWT(newClaims.UserInfo.Username)
					if err != nil {
						global.GEA_LOG.Error("get redis jwt failed", zap.Any("err", err))
					} else { // 当之前的取成功时才进行拉黑操作
						_ = service.JsonInBlacklist(model.JwtBlacklist{Jwt: RedisJwtToken})
					}
					// 无论如何都要记录当前的活跃状态
					_ = service.SetRedisJWT(accessToken, newClaims.UserInfo.Username)
				}
				c.Set("claims", newClaims)
				c.Next()
			} else {
				response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
				c.Abort()
				return
			}
		} else {
			fmt.Println("token 授权未过期")
			// TODO: 看看到时候把这里去掉，有点浪费时间
			if err, _ = service.FindUserByUuid(claims.UserInfo.UUID.String()); err != nil {
				_ = service.JsonInBlacklist(model.JwtBlacklist{Jwt: token})
				response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
				c.Abort()
			}

			if global.GEA_CONFIG.System.UseMultipoint {
				err, RedisJwtToken := service.GetRedisJWT(claims.UserInfo.Username)
				if err != nil {
					global.GEA_LOG.Error("get redis jwt failed", zap.Any("err", err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = service.JsonInBlacklist(model.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = service.SetRedisJWT(token, claims.UserInfo.Username)
			}
			c.Set("claims", claims)
			c.Next()
		}
	}
}
