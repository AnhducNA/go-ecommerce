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
	global.Logger.Info("Info log ", zap.String("OK", "success"))
	router := InitRouter()
	router.Run(":5000")
}
