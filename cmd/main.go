package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hulhay/nagasari/controller"
	"github.com/hulhay/nagasari/lib/config"
	"github.com/hulhay/nagasari/repositories"

	auth_service "github.com/hulhay/nagasari/services/auth"
	products_service "github.com/hulhay/nagasari/services/products"
	stores_service "github.com/hulhay/nagasari/services/stores"
	token_service "github.com/hulhay/nagasari/services/token"
)

var (
	cfg *config.Config = config.NewConfig()

	repo repositories.Repository = repositories.NewRepository(cfg)

	ss stores_service.StoresService     = stores_service.NewStoresService(repo)
	ps products_service.ProductsService = products_service.NewProductsService(repo)
	ts token_service.TokenService       = token_service.NewTokenService(*cfg)
	as auth_service.AuthService         = auth_service.NewAuthService(repo, ts)

	hc controller.HealthController   = controller.NewHealthController()
	sc controller.StoresController   = controller.NewStoresController(ss)
	pc controller.ProductsController = controller.NewProductsController(ps)
	ac controller.AuthController     = controller.NewAuthController(as)
)

func main() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

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
			storesRoutes.GET("/:storeUUID/products", pc.GetProducts)
		}

		// auth
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/register", ac.Register)
			authRoutes.POST("/login", ac.Login)
		}

	}

	r.Run(cfg.HTTPServerAddress())
}
