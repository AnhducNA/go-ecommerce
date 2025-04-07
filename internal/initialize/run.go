package initialize

import (
	"fmt"

	"github.com/AnhducNA/go-ecommerce/global"
)

func Run() {
	LoadConfig()
	sql := global.Config.MySQL
	fmt.Println("Server database url: ", sql.Url)
	router := InitRouter()
	router.Run(":5000")
}
