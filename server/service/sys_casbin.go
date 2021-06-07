package service

import (
	"fmt"
	"gin-element-admin/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
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
	a, err := gormadapter.NewAdapterByDB(global.GEA_DB) // 绑定Gorm
	if err != nil {
		fmt.Printf("error: adapter: %s \n", err)
	}

	m, err := model.NewModelFromString(`
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && ParamsMatch(r.obj,p.obj) && r.act == p.act
	`)
	if err != nil {
		fmt.Printf("error: model: %s \n", err)
	}
	e, err := casbin.NewEnforcer(m, a) // 根据数据库和配置文件创建执行工具
	if err != nil {
		fmt.Printf("error: enforcer: %s \n", err)
	}
	e.AddFunction("ParamsMatch", ParamsMatchFunc) // 添加自定义规则函数
	err = e.LoadPolicy()                          // 从数据库重新加载策略
	if err != nil {
		fmt.Printf("error: LoadPolicy: %s \n", err)
	}
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
