package models

import (
	"api/src/respostas"
	"strings"
	"time"
)

type Publicacao struct {
	Id        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   uint64    `json:"AutorId,omitempty"`
	AutorNick string    `json:"autornick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadoEm  time.Time `json:"criadoEm,omitempty"`
}

// Validar dados da publicação
func (publicacao *Publicacao) ValidarDados() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}
	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return respostas.MsgError("O campo título é obrigatório")
	}
	if publicacao.Conteudo == "" {
		return respostas.MsgError("O campo conteudo é obrigatório")
	}
	if publicacao.AutorId == 0 {
		return respostas.MsgError("O campo authorId é obrigatório")
	}
	return nil
}

func (publicacao *Publicacao) AtualizarCurtidas(curtidas uint64) {
	publicacao.Curtidas = curtidas
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
