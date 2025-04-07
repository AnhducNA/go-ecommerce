package initialize

import (
	"github.com/AnhducNA/go-ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.AuthMiddleware())
	// v1 := router.Group("/v1")
	// {
	// 	v1.GET("/ping", controller.NewUserController().GetUserByEmal)
	// 	v1.POST("/ping", Pong)
	// 	v1.PUT("/ping", Pong)
	// 	v1.PATCH("/ping", Pong)
	// 	v1.DELETE("/ping", Pong)
	// }
	return router
}
