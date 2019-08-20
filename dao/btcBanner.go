package dao

import (
	"encoding/json"
)

type BtcBanner struct {
	Price              float64 `json:"price"`
	TwentyFourHourVol  float64 `json:"24hr_vol"`
	AverageBlockSize   float64 `json:"average_block_size"`
	TransactionPerDay  float64 `json:"transaction_per_day"`
	MempoolSize        float64 `json:"mempool_size"`
	MarketCap          float64 `json:"market_cap"`
	CirculatingSupply  int64   `json:"circulating_supply"`
	MaxSupply          int64   `json:"max_supply"`
	HighestPriceForDay float64 `json:"highest_price_for_day"`
	LowestPriceForDay  float64 `json:"lowest_price_for_day"`
	PercentChange24h   float64 `json:"percent_change_24h"`
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
