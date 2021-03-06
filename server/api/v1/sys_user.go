package v1

import (
	"gin-element-admin/global"
	"gin-element-admin/middlewares"
	"gin-element-admin/model"
	"gin-element-admin/model/request"
	"gin-element-admin/model/response"
	"gin-element-admin/service"
	"gin-element-admin/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
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
		UserInfo: user,
		//BufferTime: global.GEA_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                                                                         // 签名生效时间
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(global.GEA_CONFIG.JWT.AccessExpiresTime)).Unix(), // 过期时间 配置文件
			Issuer:    "LLiuHuan",                                                                                // 签名的发行者
		},
	}

	accessToken, err := j.CreateToken(claims)
	if err != nil {
		global.GEA_LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	// 多点登录
	if !global.GEA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			User:        user,
			AccessToken: accessToken,
			//RefreshToken: refreshToken,
			ExpiresAt: claims.StandardClaims.ExpiresAt,
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
			User:        user,
			AccessToken: accessToken,
			//RefreshToken: refreshToken,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
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
			User:        user,
			AccessToken: accessToken,
			//RefreshToken: refreshToken,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// Register 用户注册账号
// @Tags SysUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	var R request.Register

	if errStr, err := utils.BaseValidator(&R, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	user := &model.SysUser{Username: R.Username, NickName: R.NickName, Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
	err, userReturn := service.Register(*user)
	if err != nil {
		global.GEA_LOG.Error("注册失败", zap.Any("err", err))
		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

// ChangePassword 用户修改密码
// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func ChangePassword(c *gin.Context) {
	var user request.ChangePasswordStruct
	if errStr, err := utils.BaseValidator(&user, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	U := &model.SysUser{Username: user.Username, Password: user.Password}
	if err, _ := service.ChangePassword(U, user.NewPassword); err != nil {
		global.GEA_LOG.Error("修改失败", zap.Any("err", err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// DeleteUser 删除用户
// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	var reqId request.GetById
	if errStr, err := utils.BaseValidator(&reqId, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	jwtId := getUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	if err := service.DeleteUser(reqId.ID); err != nil {
		global.GEA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// 从Gin的Context中获取从jwt解析出来的用户ID
func getUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		global.GEA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件")
		return 0
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UserInfo.ID
	}
}

// 从Gin的Context中获取从jwt解析出来的用户UUID
func getUserUuid(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.GEA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UserInfo.UUID.String()
	}
}

// SetAuthority 设置用户权限
// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func SetAuthority(c *gin.Context) {
	var sua request.SetUserAuth
	if errStr, err := utils.BaseValidator(&sua, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}

	if err := service.SetUserAuthority(sua.UUID, sua.AuthorityId); err != nil {
		global.GEA_LOG.Error("设置失败", zap.Any("err", err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

// SetUserInfo 设置用户信息
// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /user/setUserInfo [put]
func SetUserInfo(c *gin.Context) {
	var user model.SysUser
	if errStr, err := utils.BaseValidator(&user, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	if err, ReqUser := service.SetUserInfo(user); err != nil {
		global.GEA_LOG.Error("设置失败", zap.Any("err", err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "设置成功", c)
	}
}

// GetUserList 分页获取用户列表
// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	if errStr, err := utils.BaseValidatorQuery(&pageInfo, c); err != nil {
		response.FailWithMessage(errStr, c)
		return
	}
	if err, list, total := service.GetUserInfoList(pageInfo); err != nil {
		global.GEA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
