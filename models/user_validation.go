package models

import "time"

type UserValidation struct {
	BaseMode
	Code                 string `json:"code"`
	Otp                  string `json:"otp"`
	User                 User
	EmailConfirmation    bool `json:"email-confirmation"`
	PhoneConfirmation    bool `json:"phone-confirmation"`
	PasswordConfirmation bool `json:"password-confirmation"`
	IsActive             bool `json:"isActive"`
	ValidateDate         time.Time
}
