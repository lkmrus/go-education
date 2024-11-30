package user

import (
	cfg "demo/app/internal/config"
	"demo/app/pkg/db"
	"demo/app/pkg/utils"
	"encoding/json"
	"net/http"
)

type UserRepositoryDeps struct {
	DB *db.Db
}

type UserRepository struct {
	Database *db.Db
	User     *User
}

type RoleRepositoryDeps struct {
	DB *db.Db
}

type RoleRepository struct {
	Database *db.Db
	Role     *Role
}

func (repo *UserRepository) CreateUser(user *User) (*User, error) {
	result := repo.Database.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) NewUserRepository(connection *db.Db) *UserRepository {
	return &UserRepository{
		Database: connection,
	}
}

func (repo *UserRepository) Find(email string) (*User, error) {
	var user User
	res := repo.Database.Where("email = ?", email).First(&user)
	if res.Error != nil {
		if res.Error.Error() == "record not found" {
			return &user, nil
		}
		return nil, res.Error
	}

	return &user, nil
}

func (role *RoleRepository) CreateRole(writer http.ResponseWriter, request *http.Request) {
	config := cfg.Config{}
	configData := config.Init()

	var payload CreateRoleRequest

	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		Json(writer, err.Error(), 402)
		return
	}

	// TODO remove this to common area
	dbConnection := db.NewDb(configData)

	if !checkAvailableRoles(payload.Name) {
		Json(writer, "Role not available", 400)
		return
	}

	dbConnection.FirstOrCreate(role, Role{Name: role.Role.Name})
	Json(writer, role, 201)
}

func (role *RoleRepository) attachRole(writer http.ResponseWriter, request *http.Request) {
	var payload AttachRoleRequest

	userId := request.PathValue("userId")
	payload.UserID = utils.ConvertStringToUint(userId)

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
		RoleId: role.Role.ID,
	}
	tx := dbConnection.FirstOrCreate(roleUser)
	tx.Commit()

	Json(writer, roleUser, 201)
}
