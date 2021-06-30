package service

import (
	"gin-element-admin/global"
	"gin-element-admin/model"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

// Casbin 持久化到数据库  引入自定义规则
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: Casbin
//@description: 持久化到数据库  引入自定义规则
//@return: *casbin.Enforcer
func Casbin() *casbin.Enforcer {
	a, _ := gormadapter.NewAdapterByDB(global.GEA_DB)
	e, _ := casbin.NewEnforcer(global.GEA_CONFIG.Casbin.ModelPath, a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}

// ParamsMatch 自定义规则
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: ParamsMatch
//@description: 自定义规则
//@param: fullNameKey1 string, key2 string
//@return: bool
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// ParamsMatchFunc 自定义规则函数
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: ParamsMatchFunc
//@description: 自定义规则函数
//@param: args ...interface{}
//@return: interface{}, error
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return ParamsMatch(name1, name2), nil
}

// ClearCasbin 清除匹配的权限
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: ClearCasbin
//@description: 清除匹配的权限
//@param: v int, p ...string
//@return: bool
func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success

}

// UpdateCasbinApi API更新随动
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: UpdateCasbinApi
//@description: API更新随动
//@param: oldPath string, newPath string, oldMethod string, newMethod string
//@return: error
func UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.GEA_DB.Table("casbin_rule").Model(&model.CasbinModel{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}
