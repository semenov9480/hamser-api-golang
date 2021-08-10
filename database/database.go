package database

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetMethods interface {
	GetCandels(pairId int, pairName string) (primitive.M, error)
	GetLines(pairId int, pairName, value string) (primitive.M, error)
	GetLevels(pairId int, pairName, value string) (primitive.M, error)
	GetPatterns(pairId int, pairName, value string) (primitive.M, error)
}

type Database struct {
	GetMethods
}

type GetDB struct {
	client *mongo.Client
}

func InitDB(db *mongo.Client) *Database {
	return &Database{
		GetMethods: RunGetService(db),
	}

}

func RunGetService(db *mongo.Client) *GetDB {
	return &GetDB{client: db}
}

func (fp *GetDB) GetCandels(pairId int, pairName string) (primitive.M, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fp.client.Connect(ctx)
	collection := fp.client.Database(viper.GetString("DBname")).Collection("get_klines")

	var result map[string]interface{}
	//opts := options.FindOne().SetProjection(setReturn)
	err := collection.FindOne(ctx, bson.M{"pair_id": pairId, "pair_name": pairName}).Decode(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	return result, err
}

func (fp *GetDB) GetLines(pairId int, pairName, value string) (primitive.M, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fp.client.Connect(ctx)
	collection := fp.client.Database(viper.GetString("DBname")).Collection(value + "_lines")

	var result map[string]interface{}
	//err := collection.FindOne(ctx, bson.M{"pair_id": pairId, "pair_name": pairName}).Decode(&result)
	err := collection.FindOne(ctx, bson.M{"pair_id": pairId}).Decode(&result)
	return result, err
}

func (fp *GetDB) GetLevels(pairId int, pairName, value string) (primitive.M, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fp.client.Connect(ctx)
	collection := fp.client.Database(viper.GetString("DBname")).Collection(value + "_levels")

	var result map[string]interface{}
	//err := collection.FindOne(ctx, bson.M{"pair_id": pairId, "pair_name": pairName}).Decode(&result)
	err := collection.FindOne(ctx, bson.M{"pair_id": pairId}).Decode(&result)
	return result, err
}

func (fp *GetDB) GetPatterns(pairId int, pairName, value string) (primitive.M, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fp.client.Connect(ctx)
	collection := fp.client.Database(viper.GetString("DBname")).Collection("patterns_" + value)

	var result map[string]interface{}
	//err := collection.FindOne(ctx, bson.M{"pair_id": pairId, "pair_name": pairName}).Decode(&result)
	err := collection.FindOne(ctx, bson.M{"pair_id": pairId}).Decode(&result)
	return result, err
}
