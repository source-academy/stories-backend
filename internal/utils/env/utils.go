package envutils

import (
	"os"
)

func GetEnvOrDefault(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value // Includes empty string if set
	}
	return fallback
}
