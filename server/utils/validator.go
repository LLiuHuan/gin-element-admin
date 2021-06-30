package utils

import (
	"encoding/json"
	"gin-element-admin/global"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// BaseValidator POST、DELETE、PUT等使用
func BaseValidator(obj interface{}, c *gin.Context) (string, error) {
	if err := c.ShouldBind(&obj); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			return err.Error(), err
		}
		// validator.ValidationErrors类型错误则进行翻译
		errStr, _ := json.Marshal(RemoveTopStruct(errs.Translate(global.GEA_TRANS)))
		return string(errStr), err
	}
	return "", nil
}

// BaseValidatorQuery GET使用
func BaseValidatorQuery(obj interface{}, c *gin.Context) (string, error) {
	if err := c.ShouldBindQuery(obj); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			return err.Error(), err
		}
		// validator.ValidationErrors类型错误则进行翻译
		errStr, _ := json.Marshal(RemoveTopStruct(errs.Translate(global.GEA_TRANS)))
		return string(errStr), err
	}
	return "", nil
}
