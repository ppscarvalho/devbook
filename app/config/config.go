package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// URL da API a ser consumida
	APIUrl = ""
	// Porta onde a API vai estar rodando
	Porta = 0
	// Chave para assinaar o token
	HashKey []byte
	// Chave para criptografar os dados do token
	BlockKey []byte
)

// Inicializa as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		Porta = 9001
	}

	APIUrl = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}

func EndPoint(endPoint string) string {
	return fmt.Sprintf("%s/%s", APIUrl, endPoint)
}
