package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/source-academy/stories-backend/utils/constants"
)

type Config struct {
	Environment string
	Host        string
	Port        int

	SentryDSN string
}

const (
	GO_ENV = "GO_ENV"
	HOST   = "HOST"
	PORT   = "PORT"

	SENTRY = "SENTRY_DSN"
)

func LoadFromEnvironment(envFiles ...string) (*Config, error) {
	err := godotenv.Load(envFiles...)
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
		return nil, err
	}

	config := &Config{}

	// Environment
	if os.Getenv(GO_ENV) == constants.ENV_DEVELOPMENT {
		config.Environment = constants.ENV_DEVELOPMENT
	} else {
		config.Environment = constants.ENV_PRODUCTION
	}

	// Sentry
	config.SentryDSN = os.Getenv(SENTRY)

	// Server configuration
	config.Host = os.Getenv(HOST)
	envPort := os.Getenv(PORT)
	if envPort == "" {
		config.Port = constants.DEFAULT_PORT
	} else if config.Port, err = strconv.Atoi(envPort); err != nil {
		log.Fatalf("Error invalid PORT: %s\n", envPort)
		return nil, err
	}

	return config, nil
}
