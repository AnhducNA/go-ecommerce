package controller

import (
	"github.com/AnhducNA/go-ecommerce/internal/service"
	"github.com/AnhducNA/go-ecommerce/internal/vo"
	"github.com/AnhducNA/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (userCtrl *UserController) GetUserByEmal(context *gin.Context) {
	result := userCtrl.userService.Register(context.Query("email"), context.Query("purpose"))
	response.SuccessResponse(context, result, nil)
}

func (userCtrl *UserController) Register(ctx *gin.Context) {
	var params vo.UserRegistratorRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(
			ctx,
			response.ErrorParamInvalid,
			err.Error(),
		)
		return
	}
	result := userCtrl.userService.Register(
		params.Email,
		params.Purpose,
	)
	response.SuccessResponse(ctx, result, nil)
}
