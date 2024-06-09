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
	if err := k.Unmarshal("", &cfg); err != nil {
		log.Fatalln(err)
	}
	Cfg = cfg
}
