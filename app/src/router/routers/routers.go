package routers

import (
	"net/http"
	"webapp/src/middlewares"

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
	rotas = append(rotas, routerUsuarios...)
	rotas = append(rotas, rotaHome)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			router.HandleFunc(rota.Uri,
				middlewares.Logger(middlewares.Authenticate(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.Uri, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

		fileServer := http.FileServer(http.Dir("./assets/"))
		router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	}

	return router
}
