package user

import (
	"demo/app/pkg/db"
	"github.com/gorilla/mux"
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

type AttachRoleRequest struct {
	RoleName string `json:"role"`
	UserID   uint   `json:"userId"`
}

type CreateRoleRequest struct {
	Name string `json:"name"`
}

func RoleRoute(dbConnection *db.Db) *mux.Router {
	roleRepository := RoleRepository{
		Database: dbConnection,
	}

	r := mux.NewRouter()
	router := r.PathPrefix("/role").Subrouter()
	router.HandleFunc("/", roleRepository.CreateRole).Methods("POST")
	router.HandleFunc("/user/{userId}", roleRepository.attachRole).Methods("POST")
	return router
}
