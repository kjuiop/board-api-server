package controller

import (
	"board-api-server/api/form"
	"board-api-server/internal/domain/member"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberController struct {
	service member.Service
}

func NewMemberController(service member.Service) *MemberController {
	return &MemberController{
		service: service,
	}
}

func (uc *MemberController) successResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, form.ApiResponse{
		ErrorCode: form.NoError,
		Message:   form.GetCustomMessage(form.NoError),
		Result:    data,
	})
}

func (uc *MemberController) failResponse(c *gin.Context, statusCode int, errorCode int, err error) {

	logMessage := form.GetCustomErrMessage(errorCode, err.Error())
	c.Errors = append(c.Errors, &gin.Error{
		Err:  fmt.Errorf(logMessage),
		Type: gin.ErrorTypePrivate,
	})

	c.JSON(statusCode, form.ApiResponse{
		ErrorCode: errorCode,
		Message:   form.GetCustomMessage(errorCode),
	})
}

func (uc *MemberController) SignUp(c *gin.Context) {

	req := member.SignUpRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		uc.failResponse(c, http.StatusBadRequest, form.ErrParsing, fmt.Errorf("sign up json parsing err : %v", err))
		return
	}

	userId, err := uc.service.SignUp(req)
	if err != nil {
		if errors.Is(err, form.GetCustomErr(form.ErrDuplicatedUsername)) {
			uc.failResponse(c, http.StatusBadRequest, form.ErrDuplicatedUsername, fmt.Errorf("duplicated username : %w", err))
		} else {
			uc.failResponse(c, http.StatusInternalServerError, form.ErrInternalServerError, fmt.Errorf("sign up occur err : %w", err))
		}
		return
	}

	res := member.SignUpRes{
		UserId: userId,
	}

	uc.successResponse(c, http.StatusOK, form.ApiResponse{
		ErrorCode: 0,
		Message:   "success",
		Result:    res,
	})
}
