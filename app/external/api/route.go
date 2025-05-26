package api

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/configs"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/internal/interface/restful/middleware"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/registry"
	"github.com/khanghld27/kelvin-kart-challenge-api/pkg/logger"
	"net/http"
	"time"
)

const (
	envLocal      = "local"
	envStaging    = "staging"
	envProduction = "production"
)

func Restful(config *configs.Config) *gin.Engine {
	router := gin.Default()

	router.Use(ginzap.Ginzap(logger.Instance(), time.RFC3339, true))

	// Logs all panics to error log
	//router.Use(ginzap.RecoveryWithZap(logger.Instance(), true))
	if config.Env != envProduction && config.Env != envStaging {
		router.Use(middleware.CorsMiddleware())
	}

	router.Use(middleware.AddTimeout)
	router.Use(middleware.JSONWriterMiddleware)
	router.GET("/", root)

	api := router.Group("/api")

	health := api.Group("/health")
	health.GET("/healthz", healthz)

	productHandler := registry.ProductHandler()
	txnMw := registry.TransactionMiddleware()

	productRoute := api.Group("/product")
	productRoute.Use(txnMw.StartRequest)
	productRoute.Use(txnMw.EndRequest)
	productRoute.GET("/", productHandler.GetAllProducts)
	productRoute.GET("/:id", productHandler.GetProductByID)

	return router
}

func root(ctx *gin.Context) {
	type svcInfo struct {
		JSONAPI struct {
			Version string `json:"version,omitempty"`
			Name    string `json:"name,omitempty"`
		} `json:"jsonapi"`
	}

	info := svcInfo{}
	info.JSONAPI.Version = "v1"
	info.JSONAPI.Name = "Kart API"

	ctx.JSON(http.StatusOK, info)
}

func healthz(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}
