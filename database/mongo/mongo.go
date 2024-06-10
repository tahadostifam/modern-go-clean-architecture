package database

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/tahadostifam/modern-go-clean-architecture/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mu = &sync.Mutex{}
	DB *mongo.Database
)

const (
	UsersCollection = "users"
)

func NewMongoConnection() *mongo.Database {
	if DB == nil {
		mu.Lock()
		defer mu.Unlock()
		dns := fmt.Sprintf("mongodb://%s:%s@%s:%s",
			config.Cfg.Mongo.Username,
			config.Cfg.Mongo.Password,
			config.Cfg.Mongo.Host,
			config.Cfg.Mongo.Port)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(dns))
		if err != nil {
			log.Fatalln(err)
		}
		err = client.Ping(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}
		DB = client.Database(config.Cfg.Mongo.DBname)
	}
	return DB
}
