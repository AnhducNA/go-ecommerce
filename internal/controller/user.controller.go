package controller

import (
	"github.com/AnhducNA/go-ecommerce/internal/service"
	"github.com/AnhducNA/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller => service => repository => model => db
func (userCtrl *UserController) GetUserByEmal(context *gin.Context) {
	err, userInfo := userCtrl.userService.GetUserByEmail(context.Query("email"))
	if err != nil {
		response.ErrorResponse(context, 4000, "User not found")
	}
	response.SuccessResponse(context, 2000, userInfo)

}
