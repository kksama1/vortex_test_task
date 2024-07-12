package model

type Algorithm struct {
	ID       int64 `json:"id"`
	ClientID int64 `json:"clientID"`
	VWAP     bool  `json:"vwap"`
	TWAP     bool  `json:"twap"`
	HFT      bool  `json:"hft"`
}
