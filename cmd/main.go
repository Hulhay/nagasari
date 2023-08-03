package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hulhay/nagasari/controller"
	"github.com/hulhay/nagasari/lib/config"
	"github.com/hulhay/nagasari/repositories"

	stores_service "github.com/hulhay/nagasari/services/stores"
)

var (
	cfg *config.Config = config.NewConfig()

	repo repositories.Repository = repositories.NewRepository(cfg)

	ss stores_service.StoresService = stores_service.NewStoresService(repo)

	hc controller.HealthController = controller.NewHealthController()
	sc controller.StoresController = controller.NewStoresController(ss)
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

		// stores
		storesRoutes := v1.Group("/stores")
		{
			storesRoutes.GET("", sc.GetStores)
		}
	}

	r.Run(cfg.HTTPServerAddress())
}
