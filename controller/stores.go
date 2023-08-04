package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
	ss "github.com/hulhay/nagasari/services/stores"
)

type storesController struct {
	storesService ss.StoresService
}

type StoresController interface {
	GetStores(ctx *gin.Context)
}

func NewStoresController(
	storesService ss.StoresService,
) StoresController {
	return &storesController{
		storesService: storesService,
	}
}

func (sc *storesController) GetStores(ctx *gin.Context) {
	req := models.GetStoresRequest{
		Keyword: ctx.Query("keyword"),
		Page:    ctx.Query("page"),
		Size:    ctx.Query("size"),
	}

	param, err := validateGetStoresRequest(req)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, err)
		return
	}

	resp, pagination, err := sc.storesService.GetStores(ctx, param)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, err)
		return
	}

	utils.BuildAPIResponse(ctx, http.StatusOK, "success", buildGetStoresResponse(resp), pagination)
}
