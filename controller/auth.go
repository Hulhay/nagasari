package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
	as "github.com/hulhay/nagasari/services/auth"
)

type authController struct {
	authService as.AuthService
}

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

func NewAuthController(
	authService as.AuthService,
) AuthController {
	return &authController{
		authService: authService,
	}
}

func (ac *authController) Register(ctx *gin.Context) {
	var (
		request models.RegisterRequest
		err     error
	)
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, utils.ErrMalformatRequest)
		return
	}

	err = ac.validateRegisterRequest(request)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, err)
		return
	}

	err = ac.authService.Register(ctx, request)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, err)
		return
	}

	utils.BuildAPIResponse(ctx, http.StatusOK, "success", nil, nil)
}

func (ac *authController) Login(ctx *gin.Context) {
	var (
		request models.LoginRequest
		err     error
	)

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, utils.ErrMalformatRequest)
		return
	}

	err = ac.validateLoginRequest(request)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, err)
		return
	}

	response, err := ac.authService.Login(ctx, request)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, err)
		return
	}

	utils.BuildAPIResponse(ctx, http.StatusOK, "success", response, nil)
}
