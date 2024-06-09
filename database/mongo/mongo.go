package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tahadostifam/modern-go-clean-architecture/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection()*mongo.Client{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dns := fmt.Sprintf("mongodb://%s:%s@%s:%s",
	config.Cfg.Mongo.Username,
	config.Cfg.Mongo.Password,
	config.Cfg.Mongo.Host,
	config.Cfg.Mongo.Port)
	cli, err := mongo.Connect(ctx,options.Client().ApplyURI(dns))
	if err != nil {
		log.Fatal(err)
	}
	err = cli.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
	defer func() {
		if err = cli.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	return cli
}