package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"
)

// CarregarTelaLogin carrega a tela de login
func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastro carrega a tela de cadastro
func CarregarPaginaDeCadastro(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "cadastro.html", nil)
}

func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := config.EndPoint("publicacoes")
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)

	if erro != nil {
		respostas.JSONInterface(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.Mensagem(w, response)
		return
	}

	var publicacoes []models.Publicacao

	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		respostas.JSONInterface(w, http.StatusUnprocessableEntity, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	idUsuario, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.RenderTemplate(w, "home.html", struct {
		Publicacoes []models.Publicacao
		IdUsuario   uint64
	}{
		Publicacoes: publicacoes,
		IdUsuario:   idUsuario,
	})
}
