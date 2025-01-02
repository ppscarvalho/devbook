package controllers

import (
	"net/http"
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
	utils.RenderTemplate(w, "home.html", nil)
}
