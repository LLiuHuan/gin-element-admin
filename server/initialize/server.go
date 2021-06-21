package initialize

import (
	"fmt"
	"gin-element-admin/core"
	"gin-element-admin/global"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

func init() {
	// 初始化Redis
	if global.GEA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		Redis()
	}

	// 初始化Validator翻译器
	if err := InitTrans("zh"); err != nil {
		fmt.Printf("init validator failed, err: %v\n", err)
		global.GEA_LOG.Error("init validator failed, err", zap.Error(err))
		os.Exit(0)
		return
	}
}

func RunServer() *http.Server {
	// 注册路由
	router := Routers(global.GEA_CONFIG.System.Mode)

	address := fmt.Sprintf(":%d", global.GEA_CONFIG.System.Port)
	// 运行
	s := core.InitServer(address, router)
	time.Sleep(10 * time.Microsecond)

	global.GEA_LOG.Info("服务启动成功，监听端口为：  ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 Gin-Element-Admin
	当前版本:V0.0.1
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	`, address)
	fmt.Println()

	//global.GEA_LOG.Error(s.ListenAndServe().Error())
	//fmt.Println("33333333333333333")
	Close(s)
	return s
}
