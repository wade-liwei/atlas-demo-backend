package controller

import (
	"github.com/gorilla/mux"
	"github.com/wade-liwei/atlas-demo-backend/dao"
	"net/http"
)

func registerQueryBtcBanner(r *mux.Router) {
	r.HandleFunc("/btc/banner", queryBtcBanner)
}

func queryBtcBanner(w http.ResponseWriter, r *http.Request) {

	jsonAsByte, err := dao.BtcBanner{}.Get()

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(jsonAsByte)
}
