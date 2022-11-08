package repositories

import (
	"github.com/erviangelar/go-user-api/models"
	"github.com/gofrs/uuid"
)

func (p *repo) Get(id *uuid.UUID) ([]*models.User, error) {
	var l []*models.User
	err := p.db.Find(&l).Error
	return l, err
}

func (p *repo) Find(id *uuid.UUID) (*models.User, error) {
	blog := &models.User{}
	err := p.db.Where(`id = ?`, id).First(blog).Error
	return blog, err
}

func (p *repo) Create(user *models.User) (*models.User, error) {
	err := p.db.Create(user).Error
	return user, err
}
