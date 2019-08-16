//go:generate statik -src=./swagger-ui
package swagger

import (
	"github.com/gorilla/mux"

	_ "github.com/wade-liwei/atlas-demo-backend/controller/swagger/statik"

	"github.com/rakyll/statik/fs"
	"net/http"
)

func RegisterSwaggerUI(r *mux.Router) error {

	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	staticServer := http.FileServer(statikFS)
	r.PathPrefix("/swagger-ui").Handler(http.StripPrefix("/swagger-ui", staticServer))
	return nil
}
