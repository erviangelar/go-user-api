package models

type Permission struct {
	BaseMode
	Name     string `json:"name"`
	IsActive string `json:"isActive"`
}
