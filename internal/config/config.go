package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	dbutils "github.com/source-academy/stories-backend/internal/utils/db"
	envutils "github.com/source-academy/stories-backend/internal/utils/env"
)

type Config struct {
	Environment string
	Host        string
	Port        int

	Database *DatabaseConfig

	SentryDSN string

	JWKSEndpoint string
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

	JWKS_ENDPOINT = "JWKS_ENDPOINT"

	missingRequiredEnvVarMsg = "Missing required environment variable: %s"
)

func LoadFromEnvironment(envFiles ...string) (*Config, error) {
	var err error
	loadEnvFiles(envFiles...)

	config := &Config{}

	// Environment
	switch config.Environment = os.Getenv(GO_ENV); config.Environment {
	case envutils.ENV_DEVELOPMENT:
		logrus.Infoln("Running in development mode")
	case envutils.ENV_TEST:
		logrus.Infoln("Running in test mode")
	case envutils.ENV_SCRIPT:
		logrus.Infoln("Running in script mode")
	default:
		logrus.Infoln("Running in production mode")
		config.Environment = envutils.ENV_PRODUCTION
	}

	// Database
	dbConfig := &DatabaseConfig{
		// Port handled below
		TimeZone:     getEnvOrDefault(DB_TIMEZONE, dbutils.DB_DEFAULT_TIMEZONE),
		Host:         getEnvOrDefault(DB_HOSTNAME, dbutils.DB_DEFAULT_HOSTNAME),
		User:         getEnvOrDefault(DB_USERNAME, dbutils.DB_DEFAULT_USER),
		Password:     getEnvOrDefault(DB_PASSWORD, dbutils.DB_DEFAULT_PASSWORD),
		DatabaseName: getEnvOrDefault(DB_NAME, dbutils.DB_DEFAULT_NAME),
	}
	dbConfig.Port, err = parseIntFromEnv(DB_PORT, dbutils.DB_DEFAULT_PORT)
	if err != nil {
		logrus.Warningln("Invalid database port:", err)
		logrus.Warningln("Using default database port:", dbutils.DB_DEFAULT_PORT)
	}
	config.Database = dbConfig

	// Sentry
	config.SentryDSN = os.Getenv(SENTRY)

	// Server configuration
	config.Host = os.Getenv(HOST)
	config.Port, err = parseIntFromEnv(PORT, envutils.DEFAULT_PORT)
	if err != nil {
		logrus.Warningln("Invalid server port:", err)
		logrus.Warningln("Using default server port:", envutils.DEFAULT_PORT)
	}

	// RS256 public key for JWT verification
	endpoint, ok := os.LookupEnv(JWKS_ENDPOINT)
	if !ok && config.Environment != envutils.ENV_SCRIPT {
		errMsg := fmt.Sprintf(missingRequiredEnvVarMsg, JWKS_ENDPOINT)
		logrus.Errorln(errMsg)
		panic(errors.New(errMsg))
	}
	config.JWKSEndpoint = endpoint

	return config, nil
}

// Populates the environment from variables
// in the given files. If a file does not exist,
// it will be ignored (with a warning).
func loadEnvFiles(envFiles ...string) {
	// We manually iterate through each file path
	// because godotenv.Load() will stop at the first
	// error it encounters.
	for _, filePath := range envFiles {
		err := godotenv.Load(filePath)
		if err != nil {
			logrus.Warningf("Error loading env file %s: %v. Skipping...\n", filePath, err)
			continue
		}
		logrus.Infoln("Loaded environment from", filePath)
	}
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
