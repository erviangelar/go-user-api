package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/erviangelar/go-user-api/models"
	"github.com/gin-gonic/gin"
)

// Get User godoc
// @Summary Get User.
// @Description Get User.
// @Tags Users
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []UserResponse
// @Router /api/users/{id} [get]
// @Security ApiKeyAuth
func (h Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	role := ctx.Value(models.AppState{}).(models.ApplicationState).Role
	user_id := ctx.Value(models.AppState{}).(models.ApplicationState).UserID

	if strings.ToLower(role) == "user" {
		if strconv.Itoa(user_id) != id {
			c.JSON(http.StatusNotFound, gin.H{"message": "you don't have permission acces this data"})
			c.AbortWithError(http.StatusNotFound, errors.New("you don't have permission acces this data"))
			return
		}
	}
	var user UserResponse
	err := h.DB.Raw(qSelect, id).Scan(&user).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &user})
}

var qSelect = `SELECT "id","name","username", "role","created_at" as CreatedAt FROM users WHERE "id" = $1 AND deleted_at is null`
