package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/erviangelar/go-user-api/common/helper"
	"github.com/erviangelar/go-user-api/models"
	"github.com/gin-gonic/gin"
)

type PageParam struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

// Get Users godoc
// @Summary Get Users.
// @Description Get Users.
// @Tags Users
// @Accept */*
// @Produce json
// @Success 200 {object} []UserResponse
// @Router /api/users [get]
// @Security ApiKeyAuth
func (h Handler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()
	role := ctx.Value(models.AppState{}).(models.ApplicationState).Role
	user_id := ctx.Value(models.AppState{}).(models.ApplicationState).UserID

	var req PageParam
	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	pagination := helper.Pagination(c)
	offset := (pagination.Page - 1) * pagination.Limit
	roleUser := ""
	if strings.ToLower(role) == "user" {
		roleUser = `AND role ='user' AND id=` + strconv.Itoa(user_id)
	}
	var users []UserResponse
	qSelectList, qCountList := generateQuery(roleUser)
	if err := h.DB.Raw(qSelectList, pagination.Limit, offset).Scan(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	var total int
	if err := h.DB.Raw(qCountList).Scan(&total).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &users, "page": pagination.Page, "limit": pagination.Limit, "total": total})
}

func generateQuery(roleUser string) (string, string) {

	qSelectList := `SELECT "id","name","username", "role","created_at" 
	FROM users WHERE deleted_at is null ` + roleUser + ` 
	LIMIT $1
	OFFSET $2`
	qCountList := `SELECT COUNT("id") as total 
	FROM users WHERE deleted_at is null ` + roleUser
	return qSelectList, qCountList
}
