package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
	ps "github.com/hulhay/nagasari/services/products"
)

type productsController struct {
	productsService ps.ProductsService
}

type ProductsController interface {
	GetProducts(ctx *gin.Context)
}

func NewProductsController(
	productsService ps.ProductsService,
) ProductsController {
	return &productsController{
		productsService: productsService,
	}
}

func (pc *productsController) GetProducts(ctx *gin.Context) {
	req := models.GetProductsRequest{
		StoreUUID: ctx.Param("storeUUID"),
	}

	err := validateGetMenusRequest(req)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, err)
		return
	}

	store, products, err := pc.productsService.GetProductsByStoreUUID(ctx, req.StoreUUID)
	if err != nil {
		utils.BuildAPIErrorResponse(ctx, err)
		return
	}

	utils.BuildAPIResponse(ctx, http.StatusOK, "success", buildGetMenusResponse(store, products), nil)
}
