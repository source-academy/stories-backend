package config

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/source-academy/stories-backend/utils/constants"
	"github.com/stretchr/testify/assert"
)

func setupEnvFile(t *testing.T, envKeyValues map[string]string) (string, func(*testing.T)) {
	f, err := os.CreateTemp("", ".env")
	if err != nil {
		t.Error(err)
	}
	t.Log("Created temporary environment file:", f.Name())

	err = godotenv.Write(envKeyValues, f.Name())
	if err != nil {
		t.Error(err)
	}

	return f.Name(), func(t *testing.T) {
		t.Log("Removing temporary environment file:", f.Name())
		f.Close()
		os.Remove(f.Name())
	}
}

func TestLoadFromEnvironment_AppEnvironment(t *testing.T) {
	t.Run("should return development environment when GO_ENV is development", func(t *testing.T) {
		envFile, cleanUp := setupEnvFile(t, map[string]string{})
		defer cleanUp(t)

		os.Setenv(GO_ENV, constants.ENV_DEVELOPMENT)
		conf, err := LoadFromEnvironment(envFile)
		assert.Nil(t, err)
		assert.Equal(t, constants.ENV_DEVELOPMENT, conf.Environment)
	})
	t.Run("should return production environment when GO_ENV is production", func(t *testing.T) {
		envFile, cleanUp := setupEnvFile(t, map[string]string{})
		defer cleanUp(t)

		os.Setenv(GO_ENV, constants.ENV_PRODUCTION)
		conf, err := LoadFromEnvironment(envFile)
		assert.Nil(t, err)
		assert.Equal(t, constants.ENV_PRODUCTION, conf.Environment)
	})
	t.Run("should return production environment when GO_ENV anything that is not development", func(t *testing.T) {
		envFile, cleanUp := setupEnvFile(t, map[string]string{})
		defer cleanUp(t)

		os.Setenv(GO_ENV, "anything")
		conf, err := LoadFromEnvironment(envFile)
		assert.Nil(t, err)
		assert.Equal(t, constants.ENV_PRODUCTION, conf.Environment)
	})
}

func TestLoadFromEnvironment_FileEnvironment_Host(t *testing.T) {
	t.Run("should return empty string when it is not set", func(t *testing.T) {
		envFile, cleanUp := setupEnvFile(t, map[string]string{})
		defer cleanUp(t)

		conf, err := LoadFromEnvironment(envFile)
		assert.Nil(t, err)
		assert.Equal(t, "", conf.Host)
	})
	t.Run("should return host when it is set", func(t *testing.T) {
		envFile, cleanUp := setupEnvFile(t, map[string]string{
			HOST: "localhost",
		})
		defer cleanUp(t)

		conf, err := LoadFromEnvironment(envFile)
		assert.Nil(t, err)
		assert.Equal(t, "localhost", conf.Host)
	})
}

func TestLoadFromEnvironment_FileEnvironment_Port(t *testing.T) {
	t.Run("should return default port when it is not set", func(t *testing.T) {
		envFile, cleanUp := setupEnvFile(t, map[string]string{})
		defer cleanUp(t)

		conf, err := LoadFromEnvironment(envFile)
		assert.Nil(t, err)
		assert.Equal(t, constants.DEFAULT_PORT, conf.Port)
	})
	t.Run("should return port when it is set", func(t *testing.T) {
		envFile, cleanUp := setupEnvFile(t, map[string]string{
			PORT: "1234",
		})
		defer cleanUp(t)

		conf, err := LoadFromEnvironment(envFile)
		assert.Nil(t, err)
		assert.Equal(t, 1234, conf.Port)
	})
}
