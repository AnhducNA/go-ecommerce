package initialize

import (
	"github.com/AnhducNA/go-ecommerce/global"
	"github.com/AnhducNA/go-ecommerce/internal/middlewares"
	"github.com/AnhducNA/go-ecommerce/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var router *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
		router.Use(gin.Recovery())
		router.Use(gin.Logger())
	}

	router.Use(middlewares.AuthMiddleware())

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := router.Group("api/v1")
	{
		MainGroup.GET("checkStatus", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
	}
	return router
}
