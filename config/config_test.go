package config

import (
	"os"
	"testing"

	"github.com/source-academy/stories-backend/utils/constants"
	"github.com/stretchr/testify/assert"
)

func TestLoadFromEnvironment_AppEnvironment(t *testing.T) {
	t.Run("should return development environment when GO_ENV is development", func(t *testing.T) {
		os.Setenv(GO_ENV, constants.ENV_DEVELOPMENT)
		conf, err := LoadFromEnvironment()
		assert.Nil(t, err)
		assert.Equal(t, constants.ENV_DEVELOPMENT, conf.Environment)
	})
	t.Run("should return production environment when GO_ENV is production", func(t *testing.T) {
		os.Setenv(GO_ENV, constants.ENV_PRODUCTION)
		conf, err := LoadFromEnvironment()
		assert.Nil(t, err)
		assert.Equal(t, constants.ENV_PRODUCTION, conf.Environment)
	})
	t.Run("should return production environment when GO_ENV anything that is not development", func(t *testing.T) {
		os.Setenv(GO_ENV, "anything")
		conf, err := LoadFromEnvironment()
		assert.Nil(t, err)
		assert.Equal(t, constants.ENV_PRODUCTION, conf.Environment)
	})
}
