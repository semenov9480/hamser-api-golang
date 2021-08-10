package structs

type GetGraphParams struct {
	PairId   int    `form:"pairId"`
	PairName string `form:"pairName"`
}

type GetToolsParams struct {
	PairId   int    `form:"pairId"`
	PairName string `form:"pairName"`
	Value    string `form:"value"`
}
