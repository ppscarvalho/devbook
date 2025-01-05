package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// CriarPublicacao is a controller that creates a new publication
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	url := config.EndPoint("publicacoes")
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)
}

// CurtirPublicacao is a controller that likes a publication
func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	idPublicacao, erro := strconv.ParseUint(params["idPublicacao"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/%d/curtir", config.EndPoint("publicacoes"), idPublicacao)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)

	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.Mensagem(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)
}

// DesCurtirPublicacao is a controller that unlikes a publication
func DesCurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	idPublicacao, erro := strconv.ParseUint(params["idPublicacao"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/%d/descurtir", config.EndPoint("publicacoes"), idPublicacao)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)

	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)
}

// CarregarPaginaDeEdicaoPublicacao is a controller that loads the publication editing page
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idPublicacao, erro := strconv.ParseUint(params["idPublicacao"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	r.ParseForm()
	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/%d", config.EndPoint("publicacoes"), idPublicacao)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(publicacao))

	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.Mensagem(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)
}

// DeletarPublicacao is a controller that deletes a publication
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	idPublicacao, erro := strconv.ParseUint(params["idPublicacao"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/%d", config.EndPoint("publicacoes"), idPublicacao)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodDelete, url, nil)

	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)
}
