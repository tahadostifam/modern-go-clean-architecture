package config

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type (
	Config struct {
		Redis  RedisConfig
		Mongo  MongoConfig
		Jwt    JwtConfig
		Server ServerConfig
		Logger LoggerConfig
	}
	MongoConfig struct {
		Host     string
		Username string
		Password string
		Port     string
		DBname   string
	}
	JwtConfig struct {
		Secret string
	}
	ServerConfig struct {
		Port string
	}
	RedisConfig struct {
		Host     string
		Username string
		Password string
		Port     string
		DBname   int
		Protocol int
	}
	LoggerConfig struct {
		LogFilePath string
		Level       string
		Encoding    string
	}
)

var Cfg = Config{}

func NewConfig() {

	k := koanf.New(".")

	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Fatal(err)
	}

	var cfg Config

	cfg.Server.Port = k.String("server.Port")
	cfg.Mongo.Host = k.String("mongo.Host")
	cfg.Mongo.Username = k.String("mongo.Username")
	cfg.Mongo.Password = k.String("mongo.Pssword")
	cfg.Mongo.DBname = k.String("mongo.DBname")
	cfg.Mongo.Port = k.String("mongo.Port")
	cfg.Jwt.Secret = k.String("jwt.Secret")
	cfg.Redis.Host = k.String("redis.Host")
	cfg.Redis.Port = k.String("redis.Port")
	cfg.Redis.DBname = k.Int("redis.DBname")
	cfg.Redis.Protocol = k.Int("redis.Protocol")
	cfg.Redis.Username = k.String("redis.Username")
	cfg.Redis.Password = k.String("redis.Pssword")
	cfg.Logger.Level = k.String("logger.Level")
	cfg.Logger.LogFilePath = k.String("logger.LogFilePath")
	cfg.Logger.Encoding = k.String("logger.Encoding")
	Cfg = cfg
}
