package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port string
}

func (cfg *Config) Init() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8082"
	}

	return &Config{
		Port: port,
	}
}
