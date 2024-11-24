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
	dbConnection.AutoMigrate(&user.RoleUser{})

}
