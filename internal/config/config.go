package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/source-academy/stories-backend/internal/utils/constants"
)

type Config struct {
	Environment string
	Host        string
	Port        int

	Database *DatabaseConfig

	SentryDSN string
}

const (
	GO_ENV = "GO_ENV"
	HOST   = "HOST"
	PORT   = "PORT"

	DB_TIMEZONE = "DB_TIMEZONE"
	DB_HOSTNAME = "DB_HOSTNAME"
	DB_PORT     = "DB_PORT"
	DB_USERNAME = "DB_USERNAME"
	DB_PASSWORD = "DB_PASSWORD"
	DB_NAME     = "DB_NAME"

	SENTRY = "SENTRY_DSN"
)

func LoadFromEnvironment(envFiles ...string) (*Config, error) {
	err := godotenv.Load(envFiles...)
	if err != nil {
		logrus.Errorln("Error loading .env file:", err)
		return nil, err
	}

	config := &Config{}

	// Environment
	if os.Getenv(GO_ENV) == constants.ENV_DEVELOPMENT {
		config.Environment = constants.ENV_DEVELOPMENT
	} else {
		config.Environment = constants.ENV_PRODUCTION
	}

	// Database
	dbConfig := &DatabaseConfig{
		// Port handled below
		TimeZone:     getEnvOrDefault(DB_TIMEZONE, constants.DB_DEFAULT_TIMEZONE),
		Host:         getEnvOrDefault(DB_HOSTNAME, constants.DB_DEFAULT_HOSTNAME),
		User:         getEnvOrDefault(DB_USERNAME, constants.DB_DEFAULT_USER),
		Password:     getEnvOrDefault(DB_PASSWORD, constants.DB_DEFAULT_PASSWORD),
		DatabaseName: getEnvOrDefault(DB_NAME, constants.DB_DEFAULT_NAME),
	}
	dbConfig.Port, err = parseIntFromEnv(DB_PORT, constants.DB_DEFAULT_PORT)
	if err != nil {
		logrus.Warningln("Invalid database port:", err)
		logrus.Warningln("Using default database port:", constants.DB_DEFAULT_PORT)
	}
	config.Database = dbConfig

	// Sentry
	config.SentryDSN = os.Getenv(SENTRY)

	// Server configuration
	config.Host = os.Getenv(HOST)
	config.Port, err = parseIntFromEnv(PORT, constants.DEFAULT_PORT)
	if err != nil {
		logrus.Warningln("Invalid server port:", err)
		logrus.Warningln("Using default server port:", constants.DEFAULT_PORT)
	}

	return config, nil
}

func getEnvOrDefault(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value // Includes empty string if set
	}
	return fallback
}

// Parses an integer from the environment variable with the given key.
// If the environment variable is not set, it returns the default
// value. If the environment variable is set but cannot be parsed as
// an integer, it returns an error as well as setting the return value
// to the default value.
func parseIntFromEnv(key string, defaultValue int) (int, error) {
	// FIXME: Hacky abstraction
	strVal := getEnvOrDefault(key, strconv.Itoa(defaultValue))
	if strVal == "" {
		return defaultValue, nil
	}
	value, err := strconv.Atoi(strVal)
	if err != nil {
		return defaultValue, err
	}
	return value, nil
}
