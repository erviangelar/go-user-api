package models

import "github.com/gofrs/uuid"

type ApplicationState struct {
	Role        []string
	UserID      uuid.UUID
	RequestPath string
}

type (
	AppState struct{} // global struct
)
