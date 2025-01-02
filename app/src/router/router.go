package router

import (
	"webapp/src/router/routers"

	"github.com/gorilla/mux"
)

// GenerateRouter returns a new router
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routers.Configurar(r)
}
