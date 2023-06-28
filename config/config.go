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

func LoadFromEnvironment() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
		return nil, err
	}

	config := &Config{}

	if os.Getenv(GO_ENV) == constants.ENV_DEVELOPMENT {
		config.Environment = constants.ENV_DEVELOPMENT
	} else {
		config.Environment = constants.ENV_PRODUCTION
	}

	config.Host = os.Getenv(HOST)
	config.Port, err = strconv.Atoi(os.Getenv(PORT))
	if err != nil {
		log.Fatalf("Error invalid PORT: %s\n", os.Getenv(PORT))
		return nil, err
	}

	return config, nil
}
