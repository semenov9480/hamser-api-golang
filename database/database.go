package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetMethods interface {
	GetGraph() []primitive.M
}

type Database struct {
	GetMethods
}

type GetDB struct {
	client *mongo.Client
	//id     int
}

func InitDB(db *mongo.Client) *Database {
	return &Database{
		GetMethods: RunGetService(db),
	}

}

func RunGetService(db *mongo.Client) *GetDB {
	return &GetDB{client: db}
}

func (fp *GetDB) GetGraph() []primitive.M {
	var episodes []bson.M
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fp.client.Connect(ctx)
	collection := fp.client.Database("technical_analysis").Collection("get_klines")
	databases, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	if err = databases.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(episodes)
	return episodes
}
