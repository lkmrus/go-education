package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port string
	Db   Db
}

type Db struct {
	Dsn string
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

	dsn := os.Getenv("DN_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}
	var db = &Db{
		Dsn: dsn,
	}

	return &Config{
		Port: port,
		Db:   *db,
	}
}
