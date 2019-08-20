package controller

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	//cors optionsGoes Below
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", corsOpts.Handler(r)))
}
