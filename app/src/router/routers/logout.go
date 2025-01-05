package routers

import (
	"net/http"
	"webapp/src/controllers"
)

var rotaLogout = Router{
	Uri:                "/logout",
	Metodo:             http.MethodGet,
	Funcao:             controllers.FazerLogout,
	RequerAutenticacao: true,
}
