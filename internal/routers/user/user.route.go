package user

import (
	"github.com/AnhducNA/go-ecommerce/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (userRouter *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := wire.InitUserRouterHandler()
	// public router
	userRouterPublic := Router.Group("/user")
	{
		// userRouterPublic.POST("/login")
		userRouterPublic.POST("/register", userController.Register)
		// userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/info")
	}
}
