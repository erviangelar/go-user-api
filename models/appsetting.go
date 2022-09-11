package models

type ApplicationState struct {
	Role        string
	UserID      int
	RequestPath string
}

type (
	AppState struct{} // global struct
)
