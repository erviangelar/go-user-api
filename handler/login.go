package handler

import (
	"errors"
	"net/http"

	"github.com/erviangelar/go-user-api/common/jwt"
	"github.com/erviangelar/go-user-api/models"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login.
// @Description Login.
// @Tags Auth
// @Accept */*
// @Produce json
// @Param request body models.AuthRequest true "Body"
// @Success 200 {object} models.AuthResponse
// @Router /api/login [post]
func (h Handler) Login(c *gin.Context) {
	body := models.AuthRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, accessToken, refreshToken, err := h.Auth(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var userRespon models.AuthResponse
	userRespon.ID = user.UID
	userRespon.Name = user.Name
	userRespon.Role = user.Role
	userRespon.Username = user.Username
	c.JSON(http.StatusOK, gin.H{"data": &userRespon, "token": accessToken, "refresh-token": refreshToken})
}

func (h Handler) Auth(data *models.AuthRequest) (*models.User, string, string, error) {
	var user models.User
	if err := h.DB.Raw(qLogin, data.Username).Scan(&user).Error; err != nil {
		// c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		// c.AbortWithError(http.StatusNotFound, err)
		return nil, "", "", err
	}
	credErr := user.CheckPassword(data.Password)
	if credErr != nil {
		// c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		// c.AbortWithError(http.StatusNotFound, credErr)
		return nil, "", "", errors.New("invalid credentials")
	}
	accessToken, err := jwt.GenerateAccessToken(&user)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to login the user. Please try again later"})
		// c.AbortWithError(http.StatusNotFound, err)
		return nil, "", "", errors.New("failed generate token, please try again later")
	}
	refreshToken, err := jwt.GenerateRefreshToken(&user)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to login the user. Please try again later"})
		// c.AbortWithError(http.StatusNotFound, err)
		return nil, "", "", errors.New("failed generate refresh token, please try again later")
	}
	return &user, accessToken, refreshToken, nil
}

var qLogin = `SELECT "id","name","username", "role", "password" FROM users WHERE "username" = $1 AND deleted_at is null;`
