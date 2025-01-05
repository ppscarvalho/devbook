package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// CarregarTelaLogin carrega a tela de login
func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookie(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

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

func CarregarPaginaDeEdicaoPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idPublicacao, erro := strconv.ParseUint(params["idPublicacao"], 10, 64)

	if erro != nil {
		respostas.JSONInterface(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/%d", config.EndPoint("publicacoes"), idPublicacao)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)

	if erro != nil {
		respostas.JSONInterface(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.Mensagem(w, response)
		return
	}

	var publicacao models.Publicacao

	if erro = json.NewDecoder(response.Body).Decode(&publicacao); erro != nil {
		respostas.JSONInterface(w, http.StatusUnprocessableEntity, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	utils.RenderTemplate(w, "atualizar-publicacao.html", publicacao)
}

func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	url := fmt.Sprintf("%s?params=%s", config.EndPoint("usuarios"), nomeOuNick)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var usuarios []models.Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	utils.RenderTemplate(w, "usuarios.html", usuarios)
}

// CarregarPerfilDoUsuario carrega a página do perfil do usuário
func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idUsuario, erro := strconv.ParseUint(params["idUsuario"], 10, 64)

	if erro != nil {
		respostas.JSONInterface(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	usuario, erro := models.BuscarUsuarioCompleto(idUsuario, r)
	if erro != nil {
		respostas.JSONInterface(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	IdUsuarioLogado, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if idUsuario == IdUsuarioLogado {
		http.Redirect(w, r, "/perfil", http.StatusFound)
		return
	}

	utils.RenderTemplate(w, "usuario.html", struct {
		Usuario         models.Usuario
		IdUsuarioLogado uint64
	}{
		Usuario:         usuario,
		IdUsuarioLogado: IdUsuarioLogado,
	})
}

// CarregarPerfilDoUsuarioLogado carrega a página do perfil do usuário logado
func CarregarPerfilDoUsuarioLogado(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookie(r)
	idUsuario, _ := strconv.ParseUint(cookie["id"], 10, 64)

	usuario, erro := models.BuscarUsuarioCompleto(idUsuario, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	utils.RenderTemplate(w, "perfil.html", usuario)
}

// CarregarPaginaDeEdicaoDeUsuario carrega a página para edição dos dados do usuário
func CarregarPaginaDeEdicaoDeUsuario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookie(r)
	idUsuario, _ := strconv.ParseUint(cookie["id"], 10, 64)

	canal := make(chan models.Usuario)
	go models.BuscarDadosDoUsuario(canal, idUsuario, r)
	usuario := <-canal

	if usuario.Id == 0 {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: "Erro ao buscar o usuário"})
		return
	}

	utils.RenderTemplate(w, "editar-usuario.html", usuario)
}

// CarregarPaginaDeAtualizacaoDeSenha carrega a página para atualização da senha do usuário
func CarregarPaginaDeAtualizacaoDeSenha(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "atualizar-senha.html", nil)
}
