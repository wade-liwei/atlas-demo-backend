package dao

// import (
// 	"encoding/json"
// )

type WalletUser struct{}

const (
	WalletUserKey = "walletUser"
)

func (_ WalletUser) Set(bytes []byte) error {
	// jsonAsByte, err := json.Marshal(bytes)
	// if err != nil {
	// 	return err
	// }
	return Set(WalletUserKey, bytes)
}

func (_ WalletUser) Get() ([]byte, error) {

	jsonAsByte, err := Get(WalletUserKey)
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
