package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErroApi struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em JSON para a requisição
func JSONInterface(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// JSON retorna uma resposta em JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// JSONReader retorna uma resposta em JSON para a requisição
func JSONReader(w http.ResponseWriter, statusCode int, r http.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	var resposta map[string]interface{}
	if erro := json.NewDecoder(r.Body).Decode(&resposta); erro != nil {
		log.Fatal(erro)
	}

	if erro := json.NewEncoder(w).Encode(resposta); erro != nil {
		log.Fatal(erro)
	}
}

// Mensagem retorna uma resposta em JSON para a requisição
func Mensagem(w http.ResponseWriter, r *http.Response) {
	var resposta map[string]interface{}
	if erro := json.NewDecoder(r.Body).Decode(&resposta); erro != nil {
		log.Fatal(erro)
	}
	JSONInterface(w, r.StatusCode, resposta)
}

// TratarStatusCodeDeErro trata as requisições com status code 400 ou superior
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroApi
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
