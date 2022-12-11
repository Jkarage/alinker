package env

import (
	"os"

	"github.com/joho/godotenv"
)

// Env reads values from a file (default is .env) as Key:Value pair.
// Returns the string value of the given key.
// Returns an error if the key doesn't exist
func Env(key string, defaultValue string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue, nil
	}
	return value, nil
}
