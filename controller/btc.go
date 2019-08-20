package controller

import (
	"github.com/gorilla/mux"
	"github.com/wade-liwei/atlas-demo-backend/dao"
	"net/http"
)

func registerQueryBtcBanner(r *mux.Router) {
	r.HandleFunc("/btc/banner", queryBtcBanner)
	r.HandleFunc("/btc/walletuser", queryBtcWalletUser)
	r.HandleFunc("/btc/transactionrate", queryBtcTransactionRate)
	r.HandleFunc("/btc/blocklist", queryBlockList)
	r.HandleFunc("/btc/transactionpage", queryTransactionPage)
}

func queryBtcBanner(w http.ResponseWriter, r *http.Request) {

	jsonAsByte, err := dao.BtcBanner{}.Get()

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(jsonAsByte)
}

func queryBtcWalletUser(w http.ResponseWriter, r *http.Request) {
	jsonAsByte, err := dao.WalletUser{}.Get()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(jsonAsByte)
}

func queryBtcTransactionRate(w http.ResponseWriter, r *http.Request) {
	jsonAsByte, err := dao.TransactionRate{}.Get()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(jsonAsByte)
}

func queryBlockList(w http.ResponseWriter, r *http.Request) {
	jsonAsByte, err := dao.BlockList{}.Get()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(jsonAsByte)
}

func queryTransactionPage(w http.ResponseWriter, r *http.Request) {

	jsonAsByte, err := dao.TransactionPage{}.Get()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(jsonAsByte)
}
