package handler

import (
	"strings"

	utilCache "github.com/erviangelar/go-user-api/common/cache"
	"github.com/erviangelar/go-user-api/common/config"
	"github.com/erviangelar/go-user-api/common/db"
	"github.com/erviangelar/go-user-api/middleware"
	"github.com/erviangelar/go-user-api/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB     *gorm.DB
	Cache  *utilCache.RedisClient
	Repo   repositories.UserRepo
	Router *gin.Engine
	Config *config.Configurations
}

func RegisterRoutes(r *gin.Engine, Dbs *db.Dbs, config *config.Configurations) {

	repo := repositories.Init(Dbs.DB)
	h := &Handler{
		DB:     Dbs.DB,
		Cache:  Dbs.Cache,
		Repo:   repo,
		Router: r,
		Config: config,
	}
	api := r.Group("/api")
	{
		api.POST("login", h.Login)
		api.POST("register", h.Register)
		rRefresh := api.Group("/refresh-token").Use(middleware.RefreshToken())
		rRefresh.GET("", h.RefreshToken)
		routes := api.Group("/users").Use(middleware.AccessToken())
		routes.GET("", h.GetUsers)
		routes.GET("/:id", h.GetUser)
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if strings.ToLower(v) == str {
			return true
		}
	}

	return false
}
