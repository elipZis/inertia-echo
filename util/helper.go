package util

import (
	"os"
)

// Get an environment variable or return the default value
func GetEnvOrDefault(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
