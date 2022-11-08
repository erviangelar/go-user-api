package handler

import (
	"net/http"

	"github.com/erviangelar/go-user-api/common/helper"
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
func (h Handler) GetPermission(c *gin.Context) {
	var req PageParam
	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	pagination := helper.Pagination(c)
	users, err := h.FindPermission(pagination)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &users})
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
func (h Handler) GetPermissionById(c *gin.Context) {
	id := c.Param("id")
	users, err := h.FindPermissionById(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &users})
}

func (h Handler) FindPermission(pagination models.Pagination) (*[]models.UserResponse, error) {
	offset := (pagination.Page - 1) * pagination.Limit
	var users []models.UserResponse
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

func (h Handler) FindPermissionById(id string) (*models.UserResponse, error) {
	var user models.UserResponse
	err := h.DB.Raw(qPermission, id).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func queryPermission() (string, string) {
	qSelectList := `SELECT "id","name","username", "role","created_at" 
	FROM users WHERE deleted_at is null  
	LIMIT $1
	OFFSET $2`
	qCountList := `SELECT COUNT("id") as total 
	FROM users WHERE deleted_at is null `
	return qSelectList, qCountList
}

var qPermission = `SELECT "id","name","username", "role","created_at" as CreatedAt FROM users WHERE "id" = $1 AND deleted_at is null`
