package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    Host  string
    Port  string
    Debug bool
}

// Load reads environment variables and returns a Config struct
func Load() Config {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, loading system environment variables.")
    }

    // Parse and return the configuration
    return Config{
        Host:  getEnv("HOST", "127.0.0.1"),    // Default to 127.0.0.1 if HOST is not set
        Port:  getEnv("PORT", "3000"),         // Default to 3000 if PORT is not set
        Debug: getEnv("DEBUG", "false") == "true",
    }
}

// getEnv retrieves a variable from the environment, or returns a default value
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
