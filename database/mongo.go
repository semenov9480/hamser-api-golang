package database

import (
	"context"
	"log"
	"time"

	//"encoding/json"

	"go.mongodb.org/mongo-driver/mongo"

	//	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host string
	Port string
}

func MongoConnect(cfg Config) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + cfg.Host + ":" + cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		return nil, err
	}
	return client, nil
}
