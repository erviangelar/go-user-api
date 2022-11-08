package models

import (
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	UID           uuid.UUID    `gorm:"primarykey"`
	Code          string       `json:"code"`
	Username      string       `json:"username" gorm:"unique"`
	Password      string       `json:"password"`
	Name          string       `json:"name"`
	Role          []string     `json:"role"`
	Permission    []Permission `gorm:"many2many:user_permission;"`
	Phone         string       `json:"phone"`
	Email         string       `json:"email"`
	PhoneVerified bool
	EmailVerified bool
	IsActive      bool
	Tenant        []UserTenant   `gorm:"foreignKey:UserUid"`
	Relation      []UserRelation `gorm:"foreignKey:UserUid"`
}

type UserResponse struct {
	Username string `json:"username"`
	Name     string `json:"Name"`
	Role     string `json:"Role"`
}

type AuthResponse struct {
	ID       uuid.UUID
	Username string   `json:"username" example:"user@gmail.com"`
	Name     string   `json:"name" example:"user@gmail.com"`
	Role     []string `json:"role" example:"user"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID       uuid.UUID
	Username string   `json:"username" example:"user@gmail.com"`
	Name     string   `json:"name" example:"user@gmail.com"`
	Role     []string `json:"role" example:"user"`
}

type RegisterRequest struct {
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required,min=12"`
	ConfirmPassword string `json:"confirmpassword"`
}

func (request *RegisterRequest) HashPassword(password string) (*User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}
	user := User{}
	user.Username = request.Username
	user.Name = request.Username
	user.Password = string(bytes)
	return &user, nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Username = user.Email
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) error {
	user.Username = user.Email
	return nil
}
