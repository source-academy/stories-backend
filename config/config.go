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
}

const (
	GO_ENV = "GO_ENV"
	HOST   = "HOST"
	PORT   = "PORT"
)

func LoadFromEnvironment(envFiles ...string) (*Config, error) {
	err := godotenv.Load(envFiles...)
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
		return nil, err
	}

	config := &Config{}

	if os.Getenv(GO_ENV) == constants.ENV_DEVELOPMENT {
		config.Environment = constants.ENV_DEVELOPMENT
	} else {
		config.Environment = constants.ENV_PRODUCTION
	}

	config.Host = os.Getenv(HOST)

	config.Port, err = parseIntFromEnv(PORT, constants.DEFAULT_PORT)
	if err != nil {
		log.Println("WARNING: invalid port:", err)
		log.Println("Using default port:", constants.DEFAULT_PORT)
	}

	return config, nil
}

func parseIntFromEnv(key string, defaultValue int) (int, error) {
	strVal := os.Getenv(key)
	if strVal == "" {
		return defaultValue, nil
	}
	value, err := strconv.Atoi(strVal)
	if err != nil {
		return defaultValue, err
	}
	return value, nil
}
