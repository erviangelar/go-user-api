package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/erviangelar/go-user-api/common/helper"
	"github.com/erviangelar/go-user-api/models"
	"github.com/gin-gonic/gin"
)

type ResetPasswordRequest struct {
	Email           string
	Code            string
	NewPassword     string
	ConfirmPassword string
}

func (h *Handler) HResetPassword(c *gin.Context) {
	request := ResetPasswordRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := h.ResetPassword(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	message := &helper.EmailMessage{}
	mailErr := helper.SendEmail(h.Config, message)
	if mailErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": mailErr.Error()})
		c.AbortWithError(http.StatusBadRequest, mailErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password Changed"})
}

func (h *Handler) ResetPassword(request *ResetPasswordRequest) error {
	userForgot := models.ForgetPassword{}
	err := h.DB.Where("username = ? and code = ?", request.Email, request.Code).First(&userForgot)
	if err != nil {
		return errors.New("validation for change password failed")
	}
	if userForgot.Expired.Before(time.Now()) {
		return errors.New("validation expired")
	}
	userForgot.IsActive = false
	h.DB.Save(userForgot)

	user := models.User{}
	h.DB.Where("username = ?", userForgot.Email).First(&user)
	userForget := models.ForgetPassword{}
	userForget.Email = user.Email
	userForget.User = user
	userForget.Expired = time.Now().Add(time.Hour * time.Duration(1))
	h.DB.Save(&userForget)
	return nil
}
