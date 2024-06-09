package redisDB

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisConnection() *redis.Client {
	//TODO : get all from configs
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := client.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}