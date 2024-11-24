package main

import (
	cfg "demo/app/internal/config"
	user "demo/app/internal/user"
	"demo/app/pkg/db"
)

func main() {
	config := cfg.Config{}
	configData := config.Init()

	dbConnection := db.NewDb(configData)
	dbConnection.AutoMigrate(&user.User{})
	dbConnection.AutoMigrate(&user.Role{})
	dbConnection.Create(&user.Role{Name: "admin"})
	dbConnection.Create(&user.Role{Name: "user"})
	dbConnection.Create(&user.Role{Name: "guest"})
	dbConnection.Create(&user.Role{Name: "manager"})
	dbConnection.AutoMigrate(&user.RoleUser{})

}
