package dao

import (
	"encoding/json"
)

type TransactionPage struct {
	Success      bool          `json:"success"`
	Paging       Paging        `json:"paging"`
	Transactions []Transaction `json:"transactions"`
	Unconfirmed  Unconfirmed   `json:"unconfirmed"`
}

type Unconfirmed struct {
	Count    int64  `json:"count"`
	Size     int64  `json:"size"`
	Fee      string `json:"fee"`
	FeeInt   int64  `json:"fee_int"`
	Value    string `json:"value"`
	ValueInt int64  `json:"value_int"`
}

type Paging struct {
	ValidSort []string `json:"valid_sort"`
	Limit     int64    `json:"limit"`
	Sort      string   `json:"sort"`
	Dir       string   `json:"dir"`
	Prev      string   `json:"prev"`
	Next      string   `json:"next"`
	PrevLink  string   `json:"prev_link"`
	NextLink  string   `json:"next_link"`
}

type Transaction struct {
	TxId            string `json:"txid"`
	Hash            string `json:"hash"`
	Block           int64  `json:"block"`
	Confirmations   int64  `json:"confirmations"`
	Version         string `json:"version"`
	LockTime        uint64 `json:"locaktime"`
	Time            int64  `json:"time"`
	FirstSeen       int64  `json:"first_seen"`
	Propagation     string `json:"propagation"`
	DoubleSpend     bool   `json:"double_spend"`
	Size            int64  `json:"size"`
	Vsize           int64  `json:"vsize"`
	InputAmount     string `json:"input_amount"`
	InputAmountInt  int64  `json:"input_amount_int"`
	OutputAmount    string `json:"output_amount"`
	OutputAmountInt int64  `json:"output_amount_int"`
	Fee             string `json:"fee"`
	FeeInt          int64  `json:"fee_int"`
	FeeSize         string `json:"fee_size"`
	Coinbase        bool   `json:"coinbase"`
	InputCount      int64  `json:"input_count"`
	OutputCount     int64  `json:"output_count"`
	TxIndex         int64  `json:"tx_index"`
	BlockIndex      int64  `json:"block_index"`
}

const (
	TransactionPageKey = "TransactionPage"
)

func (b TransactionPage) Set() error {
	jsonAsByte, err := json.Marshal(b)
	if err != nil {
		return err
	}
	return Set(TransactionPageKey, jsonAsByte)
}

func (_ TransactionPage) Get() ([]byte, error) {

	jsonAsByte, err := Get(TransactionPageKey)
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
