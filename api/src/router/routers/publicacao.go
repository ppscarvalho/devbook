package routers

import (
	"api/src/controllers"
	"net/http"
)

// Cria rotas de Publicacao
var rotaspublicacoes = []Rota{
	{
		Uri:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
}
