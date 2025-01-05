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
	{
		Uri:                "/publicacoes/{idPublicacao}/curtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{idPublicacao}/descurtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DesCurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{idPublicacao}/atualizar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{idPublicacao}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{idPublicacao}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
}
