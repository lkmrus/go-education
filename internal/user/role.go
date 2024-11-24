package user

import (
	cfg "demo/app/internal/config"
	"demo/app/pkg/db"
	"gorm.io/gorm"
)

func checkAvailableRoles(role string) bool {
	availableRoles := make(map[string]string)
	availableRoles["admin"] = "Admin"
	availableRoles["user"] = "User"
	availableRoles["guest"] = "Guest"
	availableRoles["manager"] = "Manager"

	if _, ok := availableRoles[role]; !ok {
		return false
	}
	return true
}

type Role struct {
	gorm.Model
	Name string `json:"name" validate:"string" gorm:"uniqueIndex"`
}

type RoleUser struct {
	gorm.Model
	UserId uint
	RoleId uint
}

func (role *Role) NewRole() *Role {
	config := cfg.Config{}
	configData := config.Init()

	dbConnection := db.NewDb(configData)

	dbConnection.FirstOrCreate(role, Role{Name: role.Name})
	return role
}
