package service

import (
	"github.com/semenov9480/hamser-api-golang/database"
	"github.com/semenov9480/hamser-api-golang/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetMethods interface {
	GetCandels(input structs.GetGraphParams) (primitive.M, error)
}

type Service struct {
	GetMethods
}

func InitService(db *database.Database) *Service {
	return &Service{GetMethods: RunGetService(db)}

}
