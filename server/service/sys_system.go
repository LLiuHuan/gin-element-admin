package service

import (
	"gin-element-admin/config"
	"gin-element-admin/global"
	"gin-element-admin/utils"
	"go.uber.org/zap"
)

// GetServerInfo 获取服务器信息
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error
func GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.GEA_LOG.Error("func utils.InitCPU() Failed!", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		global.GEA_LOG.Error("func utils.InitRAM() Failed!", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.GEA_LOG.Error("func utils.InitDisk() Failed!", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}

// GetSystemConfig 读取配置文件
//@author: [LLiuHuan](https://github.com/LLiuHuan)
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: err error, conf config.Server
func GetSystemConfig() (err error, conf config.Server) {
	return nil, global.GEA_CONFIG
}
