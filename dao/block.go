package dao

import (
	"encoding/json"
)

type BlockList []Block

type Block struct {
	Hash       string `json:"hash"`
	Height     int64  `json:"height"`
	PrevBlock  string `json:"prev_block"`
	NextBlock  string `json:"next_block"`
	MrklRoot   string `json:"mrkl_root"`
	TxCount    int64  `json:"tx_count"`
	Reward     string `json:"reward"`
	Fee        string `json:"fee"`
	VoutSum    string `json:"vout_sum"`
	IsMain     bool   `json:"is_main"`
	Version    string `json:"version"`
	Difficulty string `json:"difficulty"`
	Size       int64  `json:"size"`
	Bits       string `json:"bits"`
	Nonce      string `json:"nonce"`
	Time       uint64 `json:"time"`
}

const (
	BlockListKey = "BlockList"
)

func (b BlockList) Set() error {
	jsonAsByte, err := json.Marshal(b)
	if err != nil {
		return err
	}
	return Set(BlockListKey, jsonAsByte)
}

func (_ BlockList) Get() ([]byte, error) {

	jsonAsByte, err := Get(BlockListKey)
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
