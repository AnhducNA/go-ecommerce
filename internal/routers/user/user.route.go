package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (userRouter *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	// userRouterPublic := Router.Group("/user")
	// {
	// userRouterPublic.POST("/login")
	// userRouterPublic.POST("/register")
	// userRouterPublic.POST("/otp")
	// }
	// private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/info")
	}
}
