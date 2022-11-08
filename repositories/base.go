package repositories

import (
	"github.com/erviangelar/go-user-api/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserRepo interface {
	Get(id *uuid.UUID) ([]*models.User, error)
	Find(id *uuid.UUID) (*models.User, error)
	Create(user *models.User) (*models.User, error)
}

type repo struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}
