package models

import "github.com/gofrs/uuid"

type UserRelation struct {
	BaseMode
	Code      string `json:"code"`
	IsPrimary bool   `json:"isPrimary"`
	IsActive  bool   `json:"isActive"`
	UserUid   uuid.UUID
	Child     User
}
