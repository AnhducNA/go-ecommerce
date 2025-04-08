package user

import (
	"github.com/AnhducNA/go-ecommerce/internal/controller"
	"github.com/AnhducNA/go-ecommerce/internal/repo"
	"github.com/AnhducNA/go-ecommerce/internal/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (userRouter *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userRepo := repo.NewUserRepository()
	userService := service.NewUserServie(userRepo)
	userHandlerNonDependency := controller.NewUserController(userService)
	// public router
	userRouterPublic := Router.Group("/user")
	{
		// userRouterPublic.POST("/login")
		userRouterPublic.POST("/register", userHandlerNonDependency.Register)
		// userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/info")
	}
}
