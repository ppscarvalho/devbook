package routers

import (
	"net/http"
	"webapp/src/controllers"
)

// Cria rotas de Home
var rotaHome = Router{
	Uri:                "/home",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarPaginaPrincipal,
	RequerAutenticacao: true,
}
