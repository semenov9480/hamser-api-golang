package service

import (
	"github.com/semenov9480/hamser-api-golang/database"
	"github.com/semenov9480/hamser-api-golang/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetService struct {
	database *database.Database
}

func RunGetService(db *database.Database) *GetService {
	return &GetService{database: db}
}

func (s GetService) GetCandels(input structs.GetGraphParams) (primitive.M, error) {
	setReturn := bson.D{{"_id", 0}, {"pair_id", 1}, {"pair_name", 1}, {input.TypeFrame, 1}}
	result, err := s.database.GetMethods.GetCandels(input.PairId, input.PairName, setReturn)

	return result, err
}
