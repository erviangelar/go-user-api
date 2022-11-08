package handler

import (
	"net/http"

	"github.com/erviangelar/go-user-api/models"
	"github.com/gin-gonic/gin"
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
func (h Handler) GetRole(c *gin.Context) {
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
	user, err := h.CreateRole(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	var userRespon models.RegisterResponse
	userRespon.ID = user.UID
	userRespon.Name = user.Name
	userRespon.Role = user.Role
	userRespon.Username = user.Username
	c.JSON(http.StatusOK, gin.H{"data": &userRespon})
}

// Register godoc
// @Summary Register.
// @Description Register.
// @Tags Auth
// @Accept */*
// @Produce json
// @Param request body models.RegisterRequest true "Body"
// @Success 200 {object} models.RegisterResponse
// @Router /api/register [post]
func (h Handler) GetRoleById(c *gin.Context) {
	id := c.Param("id")
	users, err := h.FindRoleById(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &users})
}

func (h Handler) FindRole(pagination models.Pagination) (*[]models.RoleResponse, error) {
	offset := (pagination.Page - 1) * pagination.Limit
	var users []models.RoleResponse
	qSelectList, qCountList := queryPermission()
	if err := h.DB.Raw(qSelectList, pagination.Limit, offset).Scan(&users).Error; err != nil {
		return nil, err
	}
	var total int
	if err := h.DB.Raw(qCountList).Scan(&total).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (h Handler) FindRoleById(id string) (*models.UserResponse, error) {
	var user models.UserResponse
	err := h.DB.Raw(qRole, id).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

var qRole = `SELECT "id","name","username", "role","created_at" as CreatedAt FROM users WHERE "id" = $1 AND deleted_at is null`
