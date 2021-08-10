package service

import (
	"github.com/semenov9480/hamser-api-golang/database"
	"github.com/semenov9480/hamser-api-golang/structs"
)

type GetService struct {
	database *database.Database
}

func RunGetService(db *database.Database) *GetService {
	return &GetService{database: db}
}

func (s GetService) GetCandels(input structs.GetGraphParams) (map[string]interface{}, error) {
	//setReturn := bson.D{{"_id", 0}, {"pair_id", 1}, {"pair_name", 1}, {"JTM", 1}, {"TTM", 1}, {"OTM", 1}}
	candels, err := s.database.GetMethods.GetCandels(input.PairId, input.PairName)
	result := map[string]interface{}{
		"candels": candels,
	}
	return result, err
}

func (s GetService) GetLines(input structs.GetToolsParams) (map[string]interface{}, error) {

	if input.Value == "All" {
		resistance, err := s.database.GetMethods.GetLines(input.PairId, input.PairName, "resistance")
		if err != nil {
			return resistance, err
		}
		support, err := s.database.GetMethods.GetLines(input.PairId, input.PairName, "support")
		if err != nil {
			return support, err
		}
		result := map[string]interface{}{
			"resistance": resistance,
			"support":    support,
		}
		return result, err
	} else {
		lines, err := s.database.GetMethods.GetLines(input.PairId, input.PairName, input.Value)
		result := map[string]interface{}{
			input.Value: lines,
		}
		return result, err
	}
}

func (s GetService) GetLevels(input structs.GetToolsParams) (map[string]interface{}, error) {

	if input.Value == "All" {
		resistance, err := s.database.GetMethods.GetLevels(input.PairId, input.PairName, "resistance")
		if err != nil {
			return resistance, err
		}
		support, err := s.database.GetMethods.GetLevels(input.PairId, input.PairName, "support")
		if err != nil {
			return support, err
		}
		result := map[string]interface{}{
			"resistance": resistance,
			"support":    support,
		}
		return result, err
	} else {
		levels, err := s.database.GetMethods.GetLevels(input.PairId, input.PairName, input.Value)
		result := map[string]interface{}{
			input.Value: levels,
		}
		return result, err
	}
}

func (s GetService) GetPatterns(input structs.GetToolsParams) (map[string]interface{}, error) {

	if input.Value == "All" {
		bearish, err := s.database.GetMethods.GetPatterns(input.PairId, input.PairName, "bearish")
		if err != nil {
			return bearish, err
		}
		bullish, err := s.database.GetMethods.GetPatterns(input.PairId, input.PairName, "bullish")
		if err != nil {
			return bullish, err
		}
		result := map[string]interface{}{
			"bearish": bearish,
			"bullish": bullish,
		}
		return result, err
	} else {
		patterns, err := s.database.GetMethods.GetPatterns(input.PairId, input.PairName, input.Value)
		result := map[string]interface{}{
			input.Value: patterns,
		}
		return result, err
	}
}
