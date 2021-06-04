package v1

import (
	"gin-element-admin/global"
	"gin-element-admin/model/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

// Captcha 生成验证码
// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成算术验证码
	//driver := base64Captcha.NewDriverDigit(global.GEA_CONFIG.Captcha.ImgHeight, global.GEA_CONFIG.Captcha.ImgWidth, global.GEA_CONFIG.Captcha.KeyLong, 0.7, 80)
	driver := base64Captcha.NewDriverMath(
		global.GEA_CONFIG.Captcha.ImgHeight,
		global.GEA_CONFIG.Captcha.ImgWidth,
		global.GEA_CONFIG.Captcha.NoiseCount,
		base64Captcha.OptionShowSlimeLine,
		&color.RGBA{},
		[]string{"wqy-microhei.ttc"},
	)

	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GEA_LOG.Error("验证码获取失败!", zap.Any("err", err))
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}
