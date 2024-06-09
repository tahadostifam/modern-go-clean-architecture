package redisDB

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/tahadostifam/modern-go-clean-architecture/config"
)

func NewRedisConnection() *redis.Client {
	//TODO : get all from configs
	client := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.Host+":"+config.Cfg.Redis.Port,
		Username:  config.Cfg.Redis.Username,
		Password: config.Cfg.Redis.Password,
		DB:      config.Cfg.Redis.DBname ,
		Protocol: config.Cfg.Redis.Protocol,
	})

	err := client.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}