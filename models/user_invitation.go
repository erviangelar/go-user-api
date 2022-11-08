package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserInvitation struct {
	BaseMode
	Code         string `json:"code"`
	IsPrimary    bool   `json:"isPrimary"`
	IsActive     bool   `json:"isActive"`
	InsitutionID uuid.UUID
	EducatorID   uuid.UUID
	EducatorCode string
	Expired      time.Time
	Role         string
}
