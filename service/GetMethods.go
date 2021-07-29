package service

import (
	"github.com/gin-gonic/gin"
	"github.com/semenov9480/hamser-api-golang/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetService struct {
	database *database.Database
}

func RunGetService(db *database.Database) *GetService {
	return &GetService{database: db}
}

func (s GetService) GetGraph(c *gin.Context) []primitive.M {
	test := s.database.GetMethods.GetGraph()
	return test

}
