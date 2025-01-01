package models

import "time"

type Publicacao struct {
	Id         uint64    `json:"id,omitempty"`
	Titulo     string    `json:"titulo,omitempty"`
	Conteudo   string    `json:"conteudo,omitempty"`
	AuthorId   uint64    `json:"authorId,omitempty"`
	AuthorNick uint64    `json:"authornick,omitempty"`
	Curtidas   uint64    `json:"curtidas"`
	CriadoEm   time.Time `json:"criadoEm,omitempty"`
}
