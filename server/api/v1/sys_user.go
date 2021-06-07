package v1

import (
	"fmt"
	"gin-element-admin/global"
	"gin-element-admin/middlewares"
	"gin-element-admin/model"
	"gin-element-admin/model/request"
	"gin-element-admin/model/response"
	"gin-element-admin/service"
	"gin-element-admin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"time"
)

// Login 登录方法
// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L request.Login

	if errStr, err := utils.BaseValidator(&L, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	if store.Verify(L.CaptchaId, L.Captcha, true) {
		U := &model.SysUser{Username: L.Username, Password: L.Password}
		if err, user := service.Login(U); err != nil {
			global.GEA_LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	j := middlewares.NewJWT() // 唯一签名

	claims := request.CustomClaims{
		UserInfo:   user,
		BufferTime: global.GEA_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                                                                         // 签名生效时间
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(global.GEA_CONFIG.JWT.AccessExpiresTime)).Unix(), // 过期时间 配置文件
			Issuer:    "LLiuHuan",                                                                                // 签名的发行者
		},
	}

	accessToken, refreshToken, err := j.CreateToken(claims)
	if err != nil {
		global.GEA_LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	// 多点登录
	if !global.GEA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			User:         user,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresAt:    claims.StandardClaims.ExpiresAt,
		}, "登录成功", c)
		return
	}
	// 单点登录需要使用redis处理token
	if err, jwtStr := service.GetRedisJWT(user.Username); err == redis.Nil {
		if err := service.SetRedisJWT(accessToken, user.Username); err != nil {
			global.GEA_LOG.Error("设置登录状态失败", zap.Any("err", err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:         user,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresAt:    claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GEA_LOG.Error("设置登录状态失败", zap.Any("err", err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := service.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := service.SetRedisJWT(accessToken, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:         user,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresAt:    claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

func Test(c *gin.Context) {
	token := c.Request.Header.Get("access-token")
	j := middlewares.NewJWT()

	parse, err := j.ParseToken(token)
	if err != nil {
		return
	}
	fmt.Println(parse)
	response.OkWithData(parse, c)
}
