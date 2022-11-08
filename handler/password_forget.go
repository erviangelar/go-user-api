package handler

import (
	"net/http"
	"time"

	"github.com/erviangelar/go-user-api/common/helper"
	"github.com/erviangelar/go-user-api/models"
	"github.com/gin-gonic/gin"
)

type ForgetPasswordRequest struct {
	username string
}

func (h *Handler) HForgetPassword(c *gin.Context) {
	request := ForgetPasswordRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := h.ForgetPassword(&request)
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

func (h *Handler) ForgetPassword(request *ForgetPasswordRequest) error {
	user := models.User{}
	h.DB.Where("username = ?", request.username).First(&user)
	userForget := models.ForgetPassword{}
	userForget.Code = helper.String(9)
	userForget.Email = user.Email
	userForget.User = user
	userForget.Expired = time.Now().Add(time.Hour * time.Duration(1))
	h.DB.Save(&userForget)
	return nil
}
