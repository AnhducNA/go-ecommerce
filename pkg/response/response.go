package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(context *gin.Context, code int, data interface{}) {
	context.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(context *gin.Context, code int, message string) {
	context.JSON(http.StatusBadRequest, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
