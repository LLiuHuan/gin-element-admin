package main

import (
	"gin-element-admin/global"
	"gin-element-admin/initialize"
)

func init() {
	global.GEA_VP = initialize.Viper() // 初始化Viper
	global.GEA_LOG = initialize.Zap()  // 初始化zap日志库
	global.GEA_DB = initialize.Gorm()  // gorm连接数据库
}

func main() {
	initialize.RunServer()

	if global.GEA_DB != nil {
		db, _ := global.GEA_DB.DB()
		defer db.Close()
	}

	// 优雅关机
	//initialize.Close(s)
}
