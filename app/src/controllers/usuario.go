package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/config"
	"webapp/src/respostas"
)

func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSONInterface(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	response, erro := http.Post(config.EndPoint("usuarios"), "application/json", bytes.NewBuffer(usuario))

	if erro != nil {
		respostas.JSONInterface(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.Mensagem(w, response)
		return
	}

	respostas.Mensagem(w, response)
}
