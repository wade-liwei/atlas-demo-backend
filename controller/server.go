package controller

import (
	"github.com/gorilla/mux"
	"github.com/wade-liwei/atlas-demo-backend/controller/swagger"
	"log"
	"net/http"
)

func Start() {
	r := mux.NewRouter()
	if err := swagger.RegisterSwaggerUI(r); err != nil {
		log.Fatal(err)
	}

	registerQueryBtcBanner(r)

	// r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
