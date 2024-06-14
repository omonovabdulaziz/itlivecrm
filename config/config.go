package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Config struct {
	Postgres Postgres
}

func Load() Config {
	dotenvFilePath := cast.ToString("./.env")
	err := godotenv.Load(dotenvFilePath)
	if err != nil {
		fmt.Println(".env file not found")
	}

	cfg := Config{
		Postgres: Postgres{
			Host:     cast.ToString("POSTGRES_HOST"),
			Port:     cast.ToString("POSTGRES_PORT"),
			User:     cast.ToString("POSTGRES_USER"),
			Password: cast.ToString("POSTGRES_PASSWORD"),
			Database: cast.ToString("POSTGRES_DB"),
		},
	}
	return cfg
}
