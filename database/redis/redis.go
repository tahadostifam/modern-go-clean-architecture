package db_redis

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/tahadostifam/modern-go-clean-architecture/config"
)

var (
	singleton   = &sync.Mutex{}
	redisClient *redis.Client
)

func GetRedisInstance(config *config.RedisConfig) *redis.Client {
	if redisClient == nil {
		singleton.Lock()
		defer singleton.Unlock()

		redisClient = redis.NewClient(&redis.Options{
			Addr:     config.Host + ":" + config.Port,
			Username: config.Username,
			Password: config.Password,
			DB:       config.DBname,
			Protocol: config.Protocol,
		})
		err := redisClient.Ping(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	return redisClient
}
