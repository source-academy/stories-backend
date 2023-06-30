package config

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/source-academy/stories-backend/internal/utils/constants"
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

func TestLoadFromEnvironment_FileEnvironment(t *testing.T) {
	t.Run("should throw error when environment file not found", func(t *testing.T) {
		_, err := LoadFromEnvironment("non-existent-file")
		assert.NotNil(t, err)
	})
	t.Run("should load a valid environment file without errors", func(t *testing.T) {
		envFile, cleanUp := setupEnvFile(t, map[string]string{
			"ANYTHING":      "anything",
			"ANYTHING_ELSE": "anything else",
		})
		defer cleanUp(t)

		_, err := LoadFromEnvironment(envFile)
		assert.Nil(t, err)
	})
	t.Run("should update environment variables from environment file", func(t *testing.T) {
		envFile, cleanUp := setupEnvFile(t, map[string]string{
			"SOMETHING": "something",
			"ANOTHER":   "other",
		})
		defer cleanUp(t)

		_, err := LoadFromEnvironment(envFile)
		assert.Nil(t, err)

		assert.Equal(t, "something", os.Getenv("SOMETHING"))
		assert.Equal(t, "other", os.Getenv("ANOTHER"))
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

func TestLoadFromEnvironment_FileEnvironment_Port_Unset(t *testing.T) {
	t.Run("should return default port when it is not set", func(t *testing.T) {
		envFile, cleanUp := setupEnvFile(t, map[string]string{})
		defer cleanUp(t)

		conf, err := LoadFromEnvironment(envFile)
		assert.Nil(t, err)
		assert.Equal(t, constants.DEFAULT_PORT, conf.Port)
	})
	t.Run("should return default port when it is set to an empty string", func(t *testing.T) {
		envFile, cleanUp := setupEnvFile(t, map[string]string{
			PORT: "",
		})
		defer cleanUp(t)

		conf, err := LoadFromEnvironment(envFile)
		assert.Nil(t, err)
		assert.Equal(t, constants.DEFAULT_PORT, conf.Port)
	})
}

func TestLoadFromEnvironment_FileEnvironment_Port_Set(t *testing.T) {
	// TODO: Investigate why this test fails
	// t.Run("should return port when it is set to a valid value", func(t *testing.T) {
	// 	envFile, cleanUp := setupEnvFile(t, map[string]string{
	// 		PORT: "1234",
	// 	})
	// 	defer cleanUp(t)

	// 	conf, err := LoadFromEnvironment(envFile)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, 1234, conf.Port)
	// })
	t.Run("should return default port when port is set to an invalid value", func(t *testing.T) {
		envFile, cleanUp := setupEnvFile(t, map[string]string{
			PORT: "not-a-number",
		})
		defer cleanUp(t)

		conf, err := LoadFromEnvironment(envFile)
		assert.Nil(t, err)
		assert.Equal(t, constants.DEFAULT_PORT, conf.Port)
	})
}

func TestParseIntFromEnv(t *testing.T) {
	testKey := "TEST_KEY"
	defaultVal := 1234

	t.Run("should return default value and no error when key is not set", func(t *testing.T) {
		intVal, err := parseIntFromEnv(testKey, defaultVal)
		assert.Equal(t, defaultVal, intVal)
		assert.Nil(t, err)
	})
	t.Run("should return default value and no error when key is set to an empty string", func(t *testing.T) {
		os.Setenv(testKey, "")

		intVal, err := parseIntFromEnv(testKey, defaultVal)
		assert.Equal(t, defaultVal, intVal)
		assert.Nil(t, err)
	})
	t.Run("should return default value when parsing an invalid numeric value", func(t *testing.T) {
		os.Setenv(testKey, "not-a-number")

		intVal, _ := parseIntFromEnv(testKey, defaultVal)
		assert.Equal(t, defaultVal, intVal)
	})
	t.Run("should throw error when parsing an invalid numeric value", func(t *testing.T) {
		os.Setenv(testKey, "not-a-number")

		_, err := parseIntFromEnv(testKey, defaultVal)
		assert.NotNil(t, err)
	})
	t.Run("should return parsed value when parsing a valid numeric value", func(t *testing.T) {
		os.Setenv(testKey, "5678")

		intVal, err := parseIntFromEnv(testKey, defaultVal)
		assert.Equal(t, 5678, intVal)
		assert.Nil(t, err)
	})
}
