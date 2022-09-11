package handler

import (
	"time"

	"github.com/erviangelar/go-user-api/middleware"
	"github.com/erviangelar/go-user-api/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB     *gorm.DB
	Repo   repositories.UserRepo
	Router *gin.Engine
}

type UserResponse struct {
	ID        uint      `example:"1" format:"int64"`
	Username  string    `json:"username" example:"admin"`
	Name      string    `json:"name" example:"Admin"`
	Role      string    `json:"role" example:"admin"`
	CreatedAt time.Time `json:"created_at" example:"04/09/2022"`
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	repo := repositories.Init(db)
	h := &Handler{
		DB:     db,
		Repo:   repo,
		Router: r,
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
