package config

import (
	"os"
	"testing"

	"github.com/source-academy/stories-backend/utils/constants"
	"github.com/stretchr/testify/assert"
)

func TestLoadFromEnvironment(t *testing.T) {
	os.Setenv(GO_ENV, constants.ENV_DEVELOPMENT)
	conf, err := LoadFromEnvironment()
	assert.Nil(t, err)
	assert.Equal(t, constants.ENV_DEVELOPMENT, conf.Environment)
}
