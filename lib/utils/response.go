package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BuildAPIResponse(ctx *gin.Context, code int, message string, data interface{}) Response {
	res := Response{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		Data: data,
	}
	ctx.JSON(code, res)
	return res
}

func BuildAPIErrorResponse(ctx *gin.Context, errorMessage error) Response {
	errRes := GetError(errorMessage)
	res := Response{
		Meta: Meta{
			Code:    int(errRes.Code()),
			Message: errRes.Error(),
		},
	}
	ctx.JSON(int(errRes.Code()), res)
	return res
}

func GetError(err error) errors.Error {
	if v, ok := err.(errors.Error); ok {
		return v
	}

	return errors.New(http.StatusInternalServerError, err.Error())
}
