package routers

import (
	"github.com/AnhducNA/go-ecommerce/internal/routers/manage"
	"github.com/AnhducNA/go-ecommerce/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Product user.ProductRouter
	Manage  manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
