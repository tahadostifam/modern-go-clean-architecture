package db_mongo

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
	singleton   = &sync.Mutex{}
	mongoClient *mongo.Database
)

const (
	UsersCollection = "users"
)

func GetMongoInstance(config *config.MongoConfig) *mongo.Database {
	if mongoClient == nil {
		singleton.Lock()
		defer singleton.Unlock()

		dns := fmt.Sprintf("mongodb://%s:%s@%s:%s",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
		)

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

		mongoClient = client.Database(config.DBname)
	}

	return mongoClient
}
