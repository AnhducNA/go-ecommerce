// go:build-wireinject
package wire

import (
	"github.com/AnhducNA/go-ecommerce/internal/controller"
	"github.com/AnhducNA/go-ecommerce/internal/repo"
	"github.com/AnhducNA/go-ecommerce/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
