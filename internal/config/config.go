package config

import (
	"log"
	"todo-api/internal/utils"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	Debug    bool
	Database DatabaseConfig
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

func Load() Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading system environment variables.")
	}

	// Return the configuration structure
	return Config{
		Host:  utils.GetEnv("HOST", "127.0.0.1"),
		Port:  utils.GetEnv("PORT", "3000"),
		Debug: utils.GetEnv("DEBUG", "false") == "true",
		Database: DatabaseConfig{
			User:     utils.GetEnv("DB_USER", "postgres"),
			Password: utils.GetEnv("DB_PASSWORD", ""),
			Host:     utils.GetEnv("DB_HOST", "localhost"),
			Port:     utils.GetEnv("DB_PORT", "5432"),
			Name:     utils.GetEnv("DB_NAME", "todo_db"),
			SSLMode:  utils.GetEnv("DB_SSLMODE", "disable"),
		},
	}
}
