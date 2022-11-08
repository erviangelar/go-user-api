package handler

import (
	"errors"
	"net/http"

	"github.com/erviangelar/go-user-api/models"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type ChangePasswordRequest struct {
	OldPassword     string
	NewPassword     string
	ConfirmPassword string
}

func (h *Handler) HChangePassword(c *gin.Context) {
	ctx := c.Request.Context()
	user_id := ctx.Value(models.AppState{}).(models.ApplicationState).UserID

	request := ChangePasswordRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if request.NewPassword != request.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "confirm password not match"})
		c.AbortWithError(http.StatusBadRequest, errors.New("confirm password not match"))
		return
	}
	err := h.ChangePassword(&user_id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password Changed"})
}

func (h *Handler) ChangePassword(uid *uuid.UUID, request ChangePasswordRequest) error {
	user := models.User{}
	h.DB.First(&user)
	err := user.HashPassword(request.NewPassword)
	if err != nil {
		return err
	}
	h.DB.Save(&user)
	return nil
}
