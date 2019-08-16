package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func registerQueryBtcBanner(r *mux.Router) {
	r.HandleFunc("/btc/banner", queryBtcBanner)
}

func queryBtcBanner(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("1234567890"))
}
