package global

import (
	"github.com/AnhducNA/go-ecommerce/pkg/logger"
	"github.com/AnhducNA/go-ecommerce/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
