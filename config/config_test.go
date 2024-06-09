package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tahadostifam/modern-go-clean-architecture/config"
)

func TestNewConfig(t *testing.T) {
	config.NewConfig()
	assert.Equal(t, "8090", config.Cfg.Server.Port)
	assert.Equal(t, "localhost", config.Cfg.Mongo.Host)
	assert.Equal(t, "mysecretkey", config.Cfg.Jwt.Secret)
	assert.Equal(t, "6379", config.Cfg.Redis.Port)
	assert.Equal(t, "../logs/log.log", config.Cfg.Logger.LogFilePath)
}
