package service

type candel struct {
	data struct {
		date     int    `json:"x"`
		closes   int    `json:"close"`
		opens    int    `json:"open"`
		high     int    `json:"high"`
		low      int    `json:"low"`
		interval string `json:"interval"`
	} `json:"data"`
}

type Candels struct {
	pair_id     int    `json:"pair_id"`
	platform    string `json:"platform"`
	pair_name   string `json:"pair_name"`
	update_date string `json:"update_date"`
	JTM         struct {
		candel *candel
	} `json:"JTM"`
	TTM struct {
		candel *candel
	} `json:"TTM"`
	OTM struct {
		candel *candel
	} `json:"OTM"`
}
