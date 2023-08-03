package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hulhay/nagasari/lib/utils"
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
	keyword := ctx.Query("keyword")
	err := validateGetStoresRequest(keyword)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, err)
		return
	}

	resp, err := sc.storesService.GetStores(ctx, keyword)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, err)
		return
	}

	utils.BuildAPIResponse(ctx, http.StatusOK, "success", buildGetStoresResponse(resp))
}
