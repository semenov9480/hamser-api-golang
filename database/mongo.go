package database

import (
	"context"
	"fmt"
	"log"
	"time"

	//"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	//	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host string
	Port string
}

type FindParams struct {
	client mongo.Client
	//id     int
}

func MongoConnect(cfg Config) (*mongo.Client, error) {
	// connection MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + cfg.Host + ":" + cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Определение типа fmt.Println(reflect.TypeOf(client))
	defer client.Disconnect(ctx)

	if err != nil {
		return nil, err
	}
	return client, nil
}

func (fp *FindParams) Find() {
	var episodes []bson.M
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fp.client.Connect(ctx)
	collection := fp.client.Database("app").Collection("get_klines")
	databases, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	if err = databases.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)

}
