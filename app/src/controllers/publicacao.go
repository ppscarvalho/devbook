package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		respostas.JSONInterface(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	url := config.EndPoint("publicacoes")
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))

	if erro != nil {
		respostas.JSONInterface(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.Mensagem(w, response)
		return
	}
	respostas.JSONInterface(w, response.StatusCode, nil)
}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar publicacoes"))
}
