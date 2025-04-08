package routers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	// v1 := router.Group("/v1")
	// {
	// 	v1.GET("/ping", controller.NewUserController.GetUserByEmal)
	// 	v1.POST("/ping", Pong)
	// 	v1.PUT("/ping", Pong)
	// 	v1.PATCH("/ping", Pong)
	// 	v1.DELETE("/ping", Pong)
	// }
	return router
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
