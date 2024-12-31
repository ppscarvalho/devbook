package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/respostas"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

// Implementar o login do usuário
func Login(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(body, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	defer db.Close()

	rep := repository.NovoRepositoryUsuario(db)
	usuarioModel, erro := rep.BuscarPorEmail(usuario.Email)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(usuarioModel.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.GenerateToken(usuarioModel.Id)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response := `{"token": "` + token + `"}`

	w.Write([]byte(response))
}
