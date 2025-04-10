package initialize

import (
	"fmt"

	"github.com/AnhducNA/go-ecommerce/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	sql := global.Config.MySQL
	fmt.Println("Server database url: ", sql.Url)
	InitLogger()
	InitMysql()
	InitRedis()
	global.Logger.Info("Info log ", zap.String("OK", "success"))
	router := InitRouter()
	router.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
