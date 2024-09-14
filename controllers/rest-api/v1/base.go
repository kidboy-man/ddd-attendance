package restapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kidboy-man/ddd-attendance/controllers/rest-api/constant"
	generic "github.com/kidboy-man/ddd-attendance/internal/generics"
)

type Controller struct{}

func (c *Controller) ReturnOK(ctx *gin.Context, httpStatus int, message string, object interface{}) {
	if message == "" {
		message = "Successful"
	}
	ctx.JSON(httpStatus, gin.H{
		"success": true,
		"code":    fmt.Sprintf("%d-0001", httpStatus),
		"message": message,
		"data":    object,
	})
}

func (c *Controller) ReturnNotOK(ctx *gin.Context, err error) {
	var message string
	var code string
	var httpStatus int

	switch err := err.(type) {
	case *generic.CustomError:
		message = err.Error()
		httpStatus = err.HTTPStatus
		code = fmt.Sprintf("%d-%d", httpStatus, err.Code)

	case validator.ValidationErrors:
		message = err.Error()
		httpStatus = http.StatusBadRequest
		code = fmt.Sprintf("%d-%d", httpStatus, constant.BadRequestErrCode)

	default:
		message = err.Error()
		httpStatus = http.StatusInternalServerError
		code = fmt.Sprintf("%d-%d", httpStatus, constant.BadRequestErrCode)
	}

	ctx.JSON(httpStatus, gin.H{
		"success": false,
		"code":    code,
		"message": message,
	})
}
