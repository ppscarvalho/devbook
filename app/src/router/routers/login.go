package routers

import (
	"net/http"
	"webapp/src/controllers"
)

var routerLogin = []Router{
	{
		Uri:                "/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaLogin,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/login",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaLogin,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/login",
		Metodo:             http.MethodPost,
		Funcao:             controllers.FazerLogin,
		RequerAutenticacao: false,
	},
}
