package initialize

import (
	"context"
	"fmt"
	"gin-element-admin/global"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() {
	// 注册路由
	router := Routers(global.GEA_CONFIG.System.Mode)

	address := fmt.Sprintf(":%d", global.GEA_CONFIG.System.Port)
	// 运行
	s := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)

	global.GEA_LOG.Info("服务启动成功，监听端口为：  ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 Gin-Element-Admin
	当前版本:V0.0.1
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	`, address)
	fmt.Println()

	global.GEA_LOG.Error(s.ListenAndServe().Error())

	// ******优雅关机******
	go func() {
		// 开启一个goroutine启动服务
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.GEA_LOG.Fatal("listen", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	global.GEA_LOG.Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := s.Shutdown(ctx); err != nil {
		global.GEA_LOG.Fatal("Server Shutdown: ", zap.Error(err))
	}

	global.GEA_LOG.Info("Server exiting")
}
