package global

import (
	"github.com/AnhducNA/go-ecommerce/pkg/logger"
	"github.com/AnhducNA/go-ecommerce/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
)
