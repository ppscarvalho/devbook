package models

import (
	"api/src/respostas"
	"api/src/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	Id       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Validar dados do usuários
func (usuario *Usuario) ValidarDados(step string) error {
	if erro := usuario.validar(step); erro != nil {
		return erro
	}

	if erro := usuario.formatar(step); erro != nil {
		return erro
	}

	return nil
}

// Valida os campos do usuário
func (usuario *Usuario) validar(step string) error {

	if usuario.Nome == "" {
		return respostas.MsgError("Campo nome é obrigatório.")
	}

	if usuario.Nick == "" {
		return respostas.MsgError("Campo nick é obrigatório.")
	}

	if usuario.Email == "" {
		return respostas.MsgError("Campo email é obrigatório.")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return respostas.MsgError("Email é inválido.")
	}

	if step == "cadastro" && usuario.Senha == "" {
		return respostas.MsgError("Campo senha é obrigatório.")
	}

	return nil
}

func (usuario *Usuario) formatar(step string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if step == "cadastro" {
		senhaComHash, erro := security.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}

	return nil
}
