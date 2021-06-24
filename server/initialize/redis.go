package initialize

import (
	"fmt"
	"gin-element-admin/global"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() {
	fmt.Println("初始化Redis")
	redisCfg := global.GEA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GEA_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.GEA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GEA_REDIS = client
	}
}
