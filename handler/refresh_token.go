package handler

import (
	"net/http"

	"github.com/erviangelar/go-user-api/common/jwt"
	"github.com/erviangelar/go-user-api/models"
	"github.com/gin-gonic/gin"
)

// Refresh Token godoc
// @Summary Refresh Token.
// @Description Refresh Token.
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} models.AuthResponse
// @Router /api/refresh-token [get]
// @Security ApiKeyAuth
func (h Handler) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	ID, _ := ctx.Value("ID").(string)

	var user models.User
	if err := h.DB.Raw(qSelectUser, ID).Scan(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	accessToken, err := jwt.GenerateAccessToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to login the user. Please try again later"})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	refreshToken, err := jwt.GenerateRefreshToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to login the user. Please try again later"})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	userRespon := models.AuthResponse{ID: user.ID, Username: user.Username, Name: user.Name, Role: user.Role}
	c.JSON(http.StatusOK, gin.H{"data": &userRespon, "token": accessToken, "refresh-token": refreshToken})
}

var qSelectUser = `SELECT "id","name","username", "role" FROM users WHERE "id" = $1 AND deleted_at is null;`
