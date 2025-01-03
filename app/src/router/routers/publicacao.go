package routers

import (
	"net/http"
	"webapp/src/controllers"
)

var routerPublicacoes = []Router{
	{
		Uri:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
}
