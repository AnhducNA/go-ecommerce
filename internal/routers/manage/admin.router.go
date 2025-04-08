package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (adminRouter *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	adminRouterPrivate := Router.Group("/admin/user")
	{
		adminRouterPrivate.POST("/active_user")
	}
}
