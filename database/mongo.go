package database

import (
	"log"
	"fmt"
	"context"
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

func MongoConnect(cfg Config) error {
	// connection MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+cfg.Host+":"+ cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
    defer client.Disconnect(ctx)
    
      /*
            List databases
    */
	var episodes []bson.M
	collection := client.Database("app").Collection("get_klines")
    databases, err := collection.Find(ctx, bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    
	if err = databases.All(ctx, &episodes); err != nil {
    	log.Fatal(err)
	}
	fmt.Println(episodes)

	
	return err
}