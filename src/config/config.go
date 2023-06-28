package config

import (
	"os"

	"github.com/source-academy/stories-backend/src/utils/constants"
)

type Config struct {
	Environment string
}

const (
	GO_ENV = "GO_ENV"
)

func LoadFromEnvironment() (*Config, error) {
	config := &Config{}

	if os.Getenv(GO_ENV) == constants.ENV_DEVELOPMENT {
		config.Environment = constants.ENV_DEVELOPMENT
	} else {
		config.Environment = constants.ENV_PRODUCTION
	}

	return config, nil
}
