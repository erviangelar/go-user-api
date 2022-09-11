package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Name     string `json:"Name"`
	Role     string `json:"Roke"`
}

type UserResponse struct {
	Username string `json:"username"`
	Name     string `json:"Name"`
	Role     string `json:"Roke"`
}

type AuthResponse struct {
	ID       uint   `example:"1" format:"int64"`
	Username string `json:"username" example:"user@gmail.com"`
	Name     string `json:"name" example:"user@gmail.com"`
	Role     string `json:"role" example:"user"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID       uint   `example:"1" format:"int64"`
	Username string `json:"username" example:"user@gmail.com"`
	Name     string `json:"name" example:"user@gmail.com"`
	Role     string `json:"role" example:"user"`
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
