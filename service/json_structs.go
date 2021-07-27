package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type candel struct {
	data struct {
		ID       primitive.ObjectID `bson: "_id"`
		date     int                `bson:"x"`
		closes   int                `bson:"close"`
		opens    int                `bson:"open"`
		high     int                `bson:"high"`
		low      int                `bson:"low"`
		interval string             `bson:"interval"`
	} `bson:"data"`
}
