package database

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetMethods interface {
	GetCandels(pairId int, pairName string, setReturn primitive.D) (primitive.M, error)
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

func (fp *GetDB) GetCandels(pairId int, pairName string, setReturn primitive.D) (primitive.M, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fp.client.Connect(ctx)
	collection := fp.client.Database(viper.GetString("DBname")).Collection("get_klines")

	var result bson.M
	opts := options.FindOne().SetProjection(setReturn)
	err := collection.FindOne(ctx, bson.M{"pair_id": pairId, "pair_name": pairName}, opts).Decode(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	return result, err
}
