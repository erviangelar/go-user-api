package models

import "github.com/gofrs/uuid"

type UserTenant struct {
	BaseMode
	Code       string `json:"code"`
	UserUid    uuid.UUID
	TenantCode string
	TenantUid  uuid.UUID
	TenantName string
	MemberUid  uuid.UUID
	MemberCode string
	MemberName string
	IsPrimary  bool `json:"isPrimary"`
	IsActive   bool `json:"isActive"`
}
