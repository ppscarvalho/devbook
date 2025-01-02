package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router representa todas as rotas da API
type Router struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(router *mux.Router) *mux.Router {
	rotas := routerLogin

	for _, rota := range rotas {
		router.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)

		// if rota.RequerAutenticacao {
		// 	router.HandleFunc(rota.Uri,
		// 		middlewares.Logger(middlewares.Authenticate(rota.Funcao)),
		// 	).Methods(rota.Metodo)
		// } else {
		// 	router.HandleFunc(rota.Uri, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		// }
	}

	return router
}
