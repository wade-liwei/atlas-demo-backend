package crontab

import (
	"encoding/json"
	"fmt"
	"github.com/wade-liwei/atlas-demo-backend/dao"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//curl -H "X-CMC_PRO_API_KEY: b5cd3d03-b51f-434a-a12a-0bdc68383c5e" -H "Accept: application/json" -G https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest
func UpdateBtcBanner() {
	fmt.Printf("run  UpdateBtcBanner\n")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	q := url.Values{}
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", ApiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request to server err: %v \n", err.Error())
		return
	}

	if strings.Compare(resp.Status, "200 OK") != 0 {
		fmt.Printf("response status:  %v \n", resp.Status)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	var m map[string]interface{}
	err = json.Unmarshal(respBody, &m)
	if err != nil {
		fmt.Printf("json unmarshal err:  %v \n", err.Error())
		return
	}

	dataArr, ok := m["data"].([]interface{})
	if !ok {
		fmt.Printf("m[\"data\"].([]interface{})  !ok")
	}

	//fmt.Println(len(dataArr))

	res := dao.BtcBanner{}
	for _, v := range dataArr {
		tmp := v.((map[string]interface{}))
		if symbol, ok := tmp["symbol"]; ok {
			if symbol.(string) == "BTC" {
				if circulatingSupply, ok := tmp["circulating_supply"].(float64); ok {
					res.CirculatingSupply = int64(circulatingSupply)
				}

				if maxSupply, ok := tmp["max_supply"].(float64); ok {
					res.MaxSupply = int64(int64(maxSupply))
				}

				if quote, ok := tmp["quote"].(map[string]interface{}); ok {
					if usd, ok := quote["USD"].(map[string]interface{}); ok {
						if price, ok := usd["price"].(float64); ok {
							res.Price = price
						}

						if volume24h, ok := usd["volume_24h"]; ok {

							if volume24hAsFloat64, ok := volume24h.(float64); ok {
								res.TwentyFourHourVol = volume24hAsFloat64
							}
						}

						if percentChange24h, ok := usd["percent_change_24h"]; ok {

							if percentChange24hAsFloat64, ok := percentChange24h.(float64); ok {
								res.PercentChange24h = percentChange24hAsFloat64
							}
						}

						if marketCap, ok := usd["market_cap"]; ok {
							if marketCapAsFloat64, ok := marketCap.(float64); ok {
								res.MarketCap = marketCapAsFloat64
							}
						}
					}
				}
				break
			}
		}
	}

	//https://api.blockchain.info/charts/n-transactions?format=json
	txPerDayJsonBytes, err := HttpGet("https://api.blockchain.info/charts/n-transactions?format=json")

	if err != nil {
		fmt.Printf("query tx per day err:  %v \n", err.Error())
		return
	}

	var txPerDay Chart

	err = json.Unmarshal(txPerDayJsonBytes, &txPerDay)
	if err != nil {
		fmt.Printf("json unmarshal err:  %v \n", err.Error())
		return
	}

	if len(txPerDay.Values) > 0 {
		res.TransactionPerDay = txPerDay.Values[len(txPerDay.Values)-1].Y
	}

	//https://blockchain.info/charts/avg-block-size?timespan=3days&format=json

	blockAverageSizeJsonBytes, err := HttpGet("https://blockchain.info/charts/avg-block-size?timespan=3days&format=json")

	if err != nil {
		fmt.Printf("query tblockAverageSizeJsonBytes err:  %v \n", err.Error())
		return
	}

	var averageBlockSize Chart

	err = json.Unmarshal(blockAverageSizeJsonBytes, &averageBlockSize)
	if err != nil {
		fmt.Printf("json unmarshal err:  %v \n", err.Error())
		return
	}

	if len(averageBlockSize.Values) > 0 {
		res.AverageBlockSize = averageBlockSize.Values[len(averageBlockSize.Values)-1].Y
	}

	//https://api.blockchain.info/charts/mempool-size?format=json
	mempoolJsonBytes, err := HttpGet("https://api.blockchain.info/charts/mempool-size?format=json")

	if err != nil {
		fmt.Printf("query mempoolJsonBytes err:  %v \n", err.Error())
		return
	}

	var mempoolChart Chart

	err = json.Unmarshal(mempoolJsonBytes, &mempoolChart)
	if err != nil {
		fmt.Printf("json unmarshal err:  %v \n", err.Error())
		return
	}

	if len(mempoolChart.Values) > 0 {
		res.MempoolSize = mempoolChart.Values[len(mempoolChart.Values)-1].Y
	}

	//curl "https://api.huobi.pro/market/detail?symbol=btcusdt"

	btcMarketSummaryAsJson, err := HttpGet("https://api.huobi.pro/market/detail?symbol=btcusdt")

	if err != nil {
		fmt.Printf("query btcMarketSummaryAsJson err:  %v \n", err.Error())
		return
	}

	var btcMarket map[string]interface{}
	err = json.Unmarshal(btcMarketSummaryAsJson, &btcMarket)
	if err != nil {
		fmt.Printf("json unmarshal err:  %v \n", err.Error())
		return
	}

	if tick, ok := btcMarket["tick"].(map[string]interface{}); ok {

		if high, ok := tick["high"].(float64); ok {
			res.HighestPriceForDay = high
		}
		if low, ok := tick["low"].(float64); ok {
			res.LowestPriceForDay = low
		}
	}

	if err := (res.Set()); err != nil {
		fmt.Println(err.Error())
	}
}

type Chart struct {
	Values []Vector `json:"values"`
}

type Vector struct {
	X int64   `json:"x"`
	Y float64 `json:"y"`
}
