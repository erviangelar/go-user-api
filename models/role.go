package models

type Role struct {
	BaseMode
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Permission  []Permission `gorm:"many2many:role_permission;"`
}

type Roles []Role

func (roles Roles) NameList() []string {
	var list []string
	for _, role := range roles {
		list = append(list, role.Name)
	}
	return list
}

type RoleResponse struct {
	BaseMode
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Permission  []Permission `json:"permissions"`
}
