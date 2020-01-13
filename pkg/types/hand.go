package types

type Hand struct {
	Uname       string   `json:"uname"`
	HandID      int      `json:"handid"`
	Balance     int      `json:"balance"`
	Hand        string   `json:"hand"`
	PotSize     int      `json:"potsize"`
	Action      string   `json:"action"`
	Amount      int      `json:"amount"`
	Board       []string `json:"board"`
	ThreeBet    bool     `json:"threebet"`
	FourPlusBet bool     `json:"fourplusbet"`
}
