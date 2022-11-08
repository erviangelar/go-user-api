package main

import (
	"log"
	"net/http"

	"github.com/erviangelar/go-user-api/app"
	"github.com/erviangelar/go-user-api/common/config"
	"github.com/erviangelar/go-user-api/common/db"
	"github.com/erviangelar/go-user-api/docs"
	"github.com/erviangelar/go-user-api/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GO User API
// @version 1.0
// @description This is a user api server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:3000
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	docs.SwaggerInfo.Title = "Go user services"
	configs := config.LoadConfig()
	dbs, err := db.Init(configs)

	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}
	//Public routes that don't have tenant checking
	excludeList := map[string]interface{}{
		"/api/v1/token":           true,
		"/health":                 true,
		"/api/v1/tenantRegister":  true,
		"/api/v1/adminLogin":      true,
		"/api/v1/clientLogin":     true,
		"/api/v1/rolePermissions": true,
		"/api/v1/authorize":       true,
		"/api/v1/tokenRefresh":    true,
		"/api/v1/oauth/token":     true,
	}

	r, err := app.InitRouter(configs, excludeList)
	if err != nil {
		log.Fatalf("Unable to initialize routes: %v\n", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")
	r.Use(cors.New(corsConfig))

	handler.RegisterRoutes(r, dbs, configs)
	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/healthcheck", HealthCheck)
	if err := r.Run(configs.Port); err != nil {
		log.Fatal(err)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Healthcheck
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /healthcheck [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}
	c.JSON(http.StatusOK, res)
}
