package controller

import (
	"board-api-server/api/form"
	"board-api-server/internal/domain/member"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	service member.Service
}

func NewUserController(service member.Service) *UserController {
	return &UserController{
		service: service,
	}
}

func (uc *UserController) successResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, form.ApiResponse{
		ErrorCode: form.NoError,
		Message:   form.GetCustomMessage(form.NoError),
		Result:    data,
	})
}

func (uc *UserController) SignUp(c *gin.Context) {
	uc.successResponse(c, http.StatusOK, form.ApiResponse{
		ErrorCode: 0,
		Message:   "success",
	})
}
