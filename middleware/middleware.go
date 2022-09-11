package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/erviangelar/go-user-api/common/jwt"
	"github.com/erviangelar/go-user-api/models"
	"github.com/gin-gonic/gin"
)

func AccessToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "you don't have an access"})
			ctx.Abort()
			return
		}
		user, err := jwt.ValidateAccessToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "access token is not valid"})
			ctx.Abort()
			return
		}
		appState := models.ApplicationState{
			Role:        user.Role,
			UserID:      int(user.ID),
			RequestPath: ctx.Request.Method,
		}
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), models.AppState{}, appState))
		ctx.Next()
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "request does not contain an refresh token"})
			ctx.Abort()
			return
		}
		ID, err := jwt.ValidateRefreshToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		IDInt, _ := strconv.Atoi(ID)
		appState := models.ApplicationState{
			Role:        "",
			UserID:      IDInt,
			RequestPath: ctx.Request.Method,
		}
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), models.AppState{}, appState))
		ctx.Next()
	}
}

func AdminValidate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "you don't have an access"})
			ctx.Abort()
			return
		}
		user, err := jwt.ValidateAccessToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		appState := models.ApplicationState{
			Role:        user.Role,
			UserID:      int(user.ID),
			RequestPath: ctx.Request.Method,
		}
		if strings.ToLower(user.Role) != "admin" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "your role don't have an access"})
			ctx.Abort()
			return
		}
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), models.AppState{}, appState))
		ctx.Next()
	}
}
