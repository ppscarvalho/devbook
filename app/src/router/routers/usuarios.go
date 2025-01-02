package routers

import (
	"net/http"
	"webapp/src/controllers"
)

var routerUsuarios = []Router{
	{
		Uri:                "/cadastrar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastro,
		RequerAutenticacao: false,
	}, {
		Uri:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CadastrarUsuario,
		RequerAutenticacao: false,
	},
}
