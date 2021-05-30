package global

import (
	"gin-element-admin/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GEA_DB     *gorm.DB
	GEA_LOG    *zap.Logger
	GEA_VP     *viper.Viper
	GEA_CONFIG config.Server
	GEA_REDIS  *redis.Client
)
