package database

import (
	"log"
	"fmt"
	"context"
	"time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host string
	Port string
}

func MongoConnect(cfg *Config) error {
	fmt.Print(cfg)

	// connection MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return err
}