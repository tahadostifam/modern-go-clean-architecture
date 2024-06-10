package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tahadostifam/modern-go-clean-architecture/config"
)

func TestNewConfig(t *testing.T) {
	config := config.Read()

	assert.NotNil(t, config)
	assert.Equal(t, "8000", config.Server.Port)
	assert.Equal(t, "localhost", config.Mongo.Host)
	assert.Equal(t, "mysecretkey", config.Jwt.Secret)
	assert.Equal(t, "6379", config.Redis.Port)
	assert.Equal(t, "../logs/log.log", config.Logger.LogFilePath)
}
