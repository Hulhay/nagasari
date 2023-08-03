package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hulhay/nagasari/controller"
	"github.com/hulhay/nagasari/lib/config"
)

var (
	cfg *config.Config = config.NewConfig()

	hc controller.HealthController = controller.NewHealthController()
)

func main() {

	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		// health
		healthRoutes := v1.Group("/health")
		{
			healthRoutes.GET("", hc.Health)
		}
	}

	r.Run(cfg.HTTPServerAddress())
}
