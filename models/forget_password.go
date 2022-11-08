package models

import "time"

type ForgetPassword struct {
	BaseMode
	Code     string    `json:"code"`
	Email    string    `json:"email"`
	User     User      `json:"user"`
	Expired  time.Time `json:"expired"`
	IsActive bool      `json:"isActive"`
}
