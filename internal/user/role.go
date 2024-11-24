package user

import (
	cfg "demo/app/internal/config"
	"demo/app/pkg/db"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
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

func (role *Role) CreateRole(writer http.ResponseWriter, request *http.Request) {
	config := cfg.Config{}
	configData := config.Init()

	var payload CreateRoleRequest

	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		Json(writer, err.Error(), 402)
		return
	}

	// TODO remove this
	dbConnection := db.NewDb(configData)

	if !checkAvailableRoles(payload.Name) {
		Json(writer, "Role not available", 400)
		return
	}

	dbConnection.FirstOrCreate(role, Role{Name: role.Name})
	Json(writer, role, 201)
}

func (role *Role) attachRole(writer http.ResponseWriter, request *http.Request) {
	var payload AttachRoleRequest

	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		Json(writer, err.Error(), 402)
		return
	}

	config := cfg.Config{}
	configData := config.Init()

	dbConnection := db.NewDb(configData)

	dbConnection.First(&role, Role{Name: payload.RoleName})

	roleUser := &RoleUser{
		UserId: payload.UserID,
		RoleId: role.ID,
	}
	tx := dbConnection.FirstOrCreate(roleUser)
	tx.Commit()

	Json(writer, roleUser, 201)
}

func RoleRoute() *mux.Router {
	role := Role{}

	r := mux.NewRouter()
	router := r.PathPrefix("/role").Subrouter()
	router.HandleFunc("/", role.CreateRole).Methods("POST")
	router.HandleFunc("/user/", role.attachRole).Methods("POST")
	return router
}
