package handler

import (
	"errors"
	"net/http"
	"net/mail"

	"github.com/erviangelar/go-user-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RoleEnum string

const (
	RoleAdmin RoleEnum = "admin"
	RoleUser  RoleEnum = "user"
)

// Register godoc
// @Summary Register.
// @Description Register.
// @Tags Auth
// @Accept */*
// @Produce json
// @Param request body models.RegisterRequest true "Body"
// @Success 200 {object} models.RegisterResponse
// @Router /api/register [post]
func (h Handler) Register(c *gin.Context) {
	body := models.RegisterRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// validate := validator.New()
	// err := validate.Struct(body)
	// if err != nil {
	// 	var errors []string
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		errors = append(errors, err.Error())
	// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "error": errors})
	// 		c.AbortWithError(http.StatusBadRequest, err)
	// 		return
	// 	}
	// }
	// if _, err := mail.ParseAddress(body.Username); err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"message": "Please provide a valid email address"})
	// 	c.AbortWithError(http.StatusNotFound, err)
	// 	return
	// }
	// if body.Password != body.ConfirmPassword {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "Password don't match"})
	// 	c.AbortWithError(http.StatusBadRequest, errors.New("password don't match"))
	// 	return
	// }
	// user, err := body.HashPassword(body.Password)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }
	// user.Role = string(RoleUser)
	// record := h.DB.Create(&user)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
	// 	c.Abort()
	// 	return
	// }
	user, err := h.Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	var userRespon models.RegisterResponse
	userRespon.ID = user.ID
	userRespon.Name = user.Name
	userRespon.Role = user.Role
	userRespon.Username = user.Username
	c.JSON(http.StatusOK, gin.H{"data": &userRespon})
}

func (h Handler) Create(data *models.RegisterRequest) (*models.User, error) {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		// var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			// errors = append(errors, err.Error())
			// c.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "error": errors})
			// c.AbortWithError(http.StatusBadRequest, err)
			return nil, err
		}
	}
	if _, err := mail.ParseAddress(data.Username); err != nil {
		// c.JSON(http.StatusNotFound, gin.H{"message": "Please provide a valid email address"})
		// c.AbortWithError(http.StatusNotFound, err)
		return nil, err
	}

	if data.Password != data.ConfirmPassword {
		// c.JSON(http.StatusBadRequest, gin.H{"message": "Password don't match"})
		// c.AbortWithError(http.StatusBadRequest, errors.New("password don't match"))
		return nil, errors.New("password don't match")
	}
	user, err := data.HashPassword(data.Password)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		// c.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}
	user.Role = string(RoleUser)
	// users, err := h.Repo.Create(user)

	record := h.DB.Create(&user)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		// c.Abort()
		return nil, err
	}
	return user, record.Error
}
