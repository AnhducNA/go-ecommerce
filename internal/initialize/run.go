package initialize

import (
	"fmt"

	"github.com/AnhducNA/go-ecommerce/global"
)

func Run() {
	LoadConfig()
	sql := global.Config.MySQL
	fmt.Println("Server database url: ", sql.Url)
	InitLogger()
	InitMysql()
	InitRedis()
	router := InitRouter()
	router.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
