package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login.html", nil)
}

func FazerLogin(w http.ResponseWriter, r *http.Request) {
	// Implementar a lógica para fazer o login do usuário
}
