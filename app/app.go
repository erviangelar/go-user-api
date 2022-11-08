package app

import (
	"github.com/erviangelar/go-user-api/common/config"
	"github.com/erviangelar/go-user-api/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter - Create gin router
func InitRouter(cfg *config.Configurations, excludeList map[string]interface{}) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))

	router.Use(middleware.TenantValidator(excludeList))
	router.Use(gin.Recovery())

	return router, nil
}
