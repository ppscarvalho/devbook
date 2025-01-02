package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Implementar a criação de uma nova publicação
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	idUsuario, erro := auth.ExtractUserId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	body, erro := io.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao

	if erro = json.Unmarshal(body, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorId = idUsuario

	if erro = publicacao.ValidarDados(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.PublicacaoRepository(db)
	publicacao.Id, erro = rep.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
}

// Implementar buscar publicações
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	idUsuario, erro := auth.ExtractUserId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.PublicacaoRepository(db)
	publicacoes, erro := rep.Buscar(idUsuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
}

// Implementar buscar uma publicação
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	idPublicacao, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.PublicacaoRepository(db)
	publicacao, erro := rep.BuscarPorId(idPublicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacao)
}

// Implementar atualizar uma publicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	idUsuario, erro := auth.ExtractUserId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)

	idPublicacao, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.PublicacaoRepository(db)
	publicacao, erro := rep.BuscarPorId(idPublicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacao.AutorId != idUsuario {
		respostas.Erro(w, http.StatusForbidden, respostas.MsgError("Usuário não autorizado, você não pode atualizar uma publicacão que não é sua."))
		return
	}

	body, erro := io.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacaoAtualizada models.Publicacao

	if erro = json.Unmarshal(body, &publicacaoAtualizada); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacaoAtualizada.Id = idPublicacao
	publicacaoAtualizada.AutorId = idUsuario

	if erro = publicacaoAtualizada.ValidarDados(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = rep.Atualizar(publicacaoAtualizada); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacaoAtualizada)
}

// Implementar deletar uma publicação
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	idUsuario, erro := auth.ExtractUserId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)

	idPublicacao, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.PublicacaoRepository(db)
	publicacao, erro := rep.BuscarPorId(idPublicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacao.AutorId != idUsuario {
		respostas.Erro(w, http.StatusForbidden, respostas.MsgError("Usuário não autorizado, você não pode deletar uma publicação que não é sua."))
		return
	}

	if erro = rep.Deletar(idPublicacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, respostas.NewResponse("Publicação deletada com sucesso"))
}

// Implementar buscar publicações por usuário tráz todas as publicações
func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idUsuario, erro := strconv.ParseUint(params["idUsuario"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.PublicacaoRepository(db)
	publicacoes, erro := rep.BuscarPorUsuario(idUsuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacoes)
}

// Implementar curtir uma publicação
func CustirPublicacao(w http.ResponseWriter, r *http.Request) {
	idUsuario, erro := auth.ExtractUserId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)

	idPublicacao, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.PublicacaoRepository(db)
	publicacao, erro := rep.BuscarPorId(idPublicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if respostas.IsEmptyStruct(publicacao) {
		respostas.Erro(w, http.StatusForbidden, respostas.MsgError("Publicação não encontrada."))
		return
	}

	if publicacao.AutorId == idUsuario {
		respostas.Erro(w, http.StatusForbidden, respostas.MsgError("Usuário não pode curtir sua própria publicação."))
		return
	}

	if erro = rep.Curtir(idPublicacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, respostas.NewResponse("Publicação curtida com sucesso."))
}

// Implementar descurtir uma publicação
func DesCustirPublicacao(w http.ResponseWriter, r *http.Request) {
	idUsuario, erro := auth.ExtractUserId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)

	idPublicacao, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.PublicacaoRepository(db)
	publicacao, erro := rep.BuscarPorId(idPublicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacao.AutorId == idUsuario {
		respostas.Erro(w, http.StatusForbidden, respostas.MsgError("Usuário não pode descurtir sua própria publicação."))
		return
	}

	if respostas.IsEmptyStruct(publicacao) {
		respostas.Erro(w, http.StatusForbidden, respostas.MsgError("Publicação não encontrada."))
		return
	}

	if erro = rep.DesCurtir(idPublicacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, respostas.NewResponse("Publicação descurtida com sucesso."))
}
