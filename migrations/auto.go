package main

import (
	user "demo/app/internal/auth"
	cfg "demo/app/internal/config"
	"demo/app/pkg/db"
)

func main() {
	config := cfg.Config{}
	configData := config.Init()

	dbConnection := db.NewDb(configData)
	dbConnection.AutoMigrate(&user.User{})
}
