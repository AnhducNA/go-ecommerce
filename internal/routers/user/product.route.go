package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (productRouter *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.POST("/detail/:id")
	}
}
