package models

// DadosAutenticacao representa os dados de autenticação do usuário.
type DadosAutenticacao struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
