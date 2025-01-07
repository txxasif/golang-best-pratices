package utils

import (
	"os"
)

// getEnv retrieves a variable from the environment, or returns a default value
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
