package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hulhay/nagasari/lib/utils"
)

type healthController struct{}

type HealthController interface {
	Health(ctx *gin.Context)
}

func NewHealthController() HealthController {
	return &healthController{}
}

func (c *healthController) Health(ctx *gin.Context) {
	utils.BuildAPIResponse(ctx, http.StatusOK, "I'm feeling fine", nil, nil)
}
