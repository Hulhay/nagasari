package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
	ms "github.com/hulhay/nagasari/services/menus"
)

type menusController struct {
	menusService ms.MenusService
}

type MenusController interface {
	GetMenus(ctx *gin.Context)
}

func NewMenusController(
	menusService ms.MenusService,
) MenusController {
	return &menusController{
		menusService: menusService,
	}
}

func (mc *menusController) GetMenus(ctx *gin.Context) {
	req := models.GetMenusRequest{
		StoreUUID: ctx.Param("storeUUID"),
	}

	err := validateGetMenusRequest(req)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, err)
		return
	}

	store, menus, err := mc.menusService.GetMenusByStoreUUID(ctx, req.StoreUUID)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, err)
		return
	}

	utils.BuildAPIResponse(ctx, http.StatusOK, "success", buildGetMenusResponse(store, menus), nil)
}
