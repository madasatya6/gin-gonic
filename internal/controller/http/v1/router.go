// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/madasatya6/gin-gonic/pkg/captcha"

	// Swagger docs.
	_ "github.com/madasatya6/gin-gonic/docs"
	"github.com/madasatya6/gin-gonic/internal/usecase"
	"github.com/madasatya6/gin-gonic/pkg/logger"

	//api
	apiv1 "github.com/madasatya6/gin-gonic/internal/controller/api/v1"

	//modules
	"github.com/madasatya6/gin-gonic/internal/controller/modules/v1/example"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.Translation) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.GET("/captcha", captcha.GenerateCaptcha)

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newTranslationRoutes(h, t, l)
	}

	apiRoute := handler.Group("/api")
	{
		apiv1.Init(apiRoute, t, l)
	}

	//example route
	exmple := handler.Group("/example")
	{
		example.Init(exmple, t, l)
	}

	
}
