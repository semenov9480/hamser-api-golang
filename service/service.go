package service

import (
	"github.com/semenov9480/hamser-api-golang/database"
	"github.com/semenov9480/hamser-api-golang/structs"
)

type GetMethods interface {
	GetCandels(input structs.GetGraphParams) (map[string]interface{}, error)
	GetLines(input structs.GetToolsParams) (map[string]interface{}, error)
	GetLevels(input structs.GetToolsParams) (map[string]interface{}, error)
	GetPatterns(input structs.GetToolsParams) (map[string]interface{}, error)
}

type Service struct {
	GetMethods
}

func InitService(db *database.Database) *Service {
	return &Service{GetMethods: RunGetService(db)}

}
