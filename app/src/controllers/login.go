package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/respostas"
)

// Implementar a lógica para fazer o login do usuário
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSONInterface(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	response, erro := http.Post(config.EndPoint("login"), "application/json", bytes.NewBuffer(usuario))

	if erro != nil {
		respostas.JSONInterface(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.Mensagem(w, response)
		return
	}

	var dadosAutenticacao models.DadosAutenticacao
	if erro = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); erro != nil {
		respostas.JSONInterface(w, http.StatusUnprocessableEntity, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	if erro = cookies.Salvar(w, dadosAutenticacao.Id, dadosAutenticacao.Token); erro != nil {
		respostas.JSONInterface(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	respostas.JSONInterface(w, http.StatusOK, nil)
}
