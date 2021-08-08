package structs

type GetGraphParams struct {
	PairId   int    `form:"pairId"`
	PairName string `form:"pairName"`
}
