package dao

import (
	"encoding/json"
)

type BtcBanner struct {
	AverageBlockSize   float64 `json:"average_block_size"`
	TransactionPerDay  int64   `json:"transaction_per_day"`
	MempoolSize        int64   `json:"mempool_size"`
	MarketCap          int64   `json:"market_cap"`
	Volume             int64   `json:"volume"`
	CirculatingSupply  int64   `json:"circulating_supply"`
	MaxSupply          int64   `json:"max_supply"`
	MarketPrice        int64   `json:"market_price"`
	HighestPriceForDay float64 `json:"highest_price_for_day"`
	LowestPriceForDay  float64 `json:"lowest_price_for_day"`
	ChangeForDay       float64 `json:"change_for_day"`
}

const (
	BtcBannerKey = "BtcBanner"
)

func (b BtcBanner) Set() error {
	jsonAsByte, err := json.Marshal(b)
	if err != nil {
		return err
	}
	return Set(BtcBannerKey, jsonAsByte)
}

func (_ BtcBanner) Get() ([]byte, error) {

	jsonAsByte, err := Get(BtcBannerKey)
	if err != nil {
		return nil, err
	}
	return jsonAsByte, nil
	// res := BtcBanner{}
	//
	// if err := json.Unmarshal(jsonAsByte, &res); err != nil {
	// 	return nil, err
	// }
	//
	// return &res, nil
}
