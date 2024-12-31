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
	"strings"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(body, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.ValidarDados("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	rep := repository.NovoRepositoryUsuario(db)
	usuario.Id, erro = rep.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuario, _ = rep.BuscarPorId(usuario.Id)

	respostas.JSON(w, http.StatusCreated, usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	params := strings.ToLower(r.URL.Query().Get("params"))

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.NovoRepositoryUsuario(db)
	usuarios, erro := rep.Buscar(params)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	usuarioId, erro := strconv.ParseUint(param["id"], 10, 64)

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

	rep := repository.NovoRepositoryUsuario(db)
	usuarios, erro := rep.BuscarPorId(usuarioId)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	usuarioId, erro := strconv.ParseUint(param["id"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var resp = validaUsuarioToken(w, r, usuarioId)
	if !resp {
		return
	}

	payload, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(payload, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.ValidarDados("alteracao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.NovoRepositoryUsuario(db)
	usuarios, erro := rep.Atualizar(usuarioId, usuario)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	usuarioId, erro := strconv.ParseUint(param["id"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var resp = validaUsuarioToken(w, r, usuarioId)
	if !resp {
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.NovoRepositoryUsuario(db)

	if erro := rep.Excluir(usuarioId); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	idSeguidor, erro := auth.ExtractUserId(r)

	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	idUsuario, erro := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if idUsuario == idSeguidor {
		respostas.Erro(w, http.StatusForbidden, respostas.MsgError("Não é possível seguir você mesmo"))
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.NovoRepositoryUsuario(db)

	if erro := rep.SeguirUsuario(idUsuario, idSeguidor); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	idSeguidor, erro := auth.ExtractUserId(r)

	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	idUsuario, erro := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if idUsuario == idSeguidor {
		respostas.Erro(w, http.StatusForbidden, respostas.MsgError("Não é possível parar de seguir você mesmo"))
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	rep := repository.NovoRepositoryUsuario(db)

	if erro := rep.PararDeSeguirUsuario(idUsuario, idSeguidor); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {

	idUsuario, erro := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
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

	rep := repository.NovoRepositoryUsuario(db)
	usuarios, erro := rep.BuscarSeguidores(idUsuario)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {

	idSeguidor, erro := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
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

	rep := repository.NovoRepositoryUsuario(db)
	usuarios, erro := rep.BuscarSeguindo(idSeguidor)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

func validaUsuarioToken(w http.ResponseWriter, r *http.Request, usuarioId uint64) bool {
	checkUserIdInToken, erro := auth.ExtractUserId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return false
	}

	if checkUserIdInToken != usuarioId {
		respostas.Erro(w, http.StatusForbidden, respostas.MsgError("Acesso negado"))
		return false
	}
	return true
}
