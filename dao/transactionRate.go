package dao

// import (
// 	"encoding/json"
// )

type TransactionRate struct{}

const (
	TransactionRateKey = "transactionRate"
)

func (_ TransactionRate) Set(bytes []byte) error {
	// jsonAsByte, err := json.Marshal(bytes)
	// if err != nil {
	// 	return err
	// }
	return Set(TransactionRateKey, bytes)
}

func (_ TransactionRate) Get() ([]byte, error) {

	jsonAsByte, err := Get(TransactionRateKey)
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
