package service

import (
	"github.com/gin-gonic/gin"
	"github.com/semenov9480/hamser-api-golang/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetMethods interface {
	GetGraph(g *gin.Context) []primitive.M
}

type Service struct {
	GetMethods
}

func InitService(db *database.Database) *Service {
	return &Service{GetMethods: RunGetService(db)}

}
