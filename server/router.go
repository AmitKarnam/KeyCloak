package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AmitKarnam/KeyCloak/controllers"
	"github.com/AmitKarnam/KeyCloak/internal/utlis/logger/zapLogger"
)

func InitRouter() (*gin.Engine, error) {

	gin.SetMode(gin.ReleaseMode)

	// Initialise New gin engine
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%v %v %v %v", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	healthController := controllers.HealthController{}
	router.GET("/health", healthController.Status)

	service := router.Group("keycloak")
	{
		apiGroup := service.Group("api")
		{
			versionGroup := apiGroup.Group("v1")
			{
				secretEngineGroup := versionGroup.Group("secretengine")
				secretEngineGroup.GET("", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{"Secrets GET": "GET Endpoint for secret engine"})
				})

				zapLogger.KeyCloaklogger.Infof("Secret Engine Cotroller Initialized")

			}

			{
				dataGroup := versionGroup.Group("data")
				dataGroup.GET(":name", func(c *gin.Context) {
					secretName := c.Param("name")
					c.JSON(http.StatusOK, gin.H{
						"Secret Engine Name": secretName,
					})
				})

				zapLogger.KeyCloaklogger.Infof("Data Controller Initialized")
			}
		}
	}

	return router, nil
}
