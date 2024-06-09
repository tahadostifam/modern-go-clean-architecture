package redisDB

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/tahadostifam/modern-go-clean-architecture/config"
)

var (
	Mu  = &sync.Mutex{}
	RDB *redis.Client
)

func NewRedisConnection() *redis.Client {
	if RDB == nil {
		Mu.Lock()
		defer Mu.Unlock()
		RDB := redis.NewClient(&redis.Options{
			Addr:     config.Cfg.Redis.Host + ":" + config.Cfg.Redis.Port,
			Username: config.Cfg.Redis.Username,
			Password: config.Cfg.Redis.Password,
			DB:       config.Cfg.Redis.DBname,
			Protocol: config.Cfg.Redis.Protocol,
		})
		err := RDB.Ping(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
	return RDB
}
