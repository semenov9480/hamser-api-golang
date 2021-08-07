package structs

type GetGraphParams struct {
	PairId    int    `json:"pairId"`
	PairName  string `json:"pairName"`
	TypeFrame string `json:"typeFrame"`
}
