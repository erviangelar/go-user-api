package main

import (
	"log"
	"net/http"

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

	docs.SwaggerInfo.Title = "GO User Example API"

	configs := config.LoadConfig()
	db := db.Init(configs)

	r := gin.Default()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")
	r.Use(cors.New(corsConfig))

	handler.RegisterRoutes(r, db)
	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/healthcheck", HealthCheck)
	//Start Server
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
