package auth

import (
	"api/src/config"
	"api/src/respostas"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(Id uint64) (string, error) {
	permissions := jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
		"usuarioId":  Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

// Validando o token informado pelo usuário
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	if tokenString == "" {
		return respostas.MsgError("Token não fornecido")
	}

	token, erro := jwt.Parse(tokenString, returnKeyVerification)

	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return respostas.MsgError("Token inválido")
}

func extractToken(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")

	if len(strings.Split(tokenString, " ")) == 2 {
		return strings.Split(tokenString, " ")[1] // remove "Bearer " from the beginning of the token
	}

	return ""
}

// Retorna a chave de verificação do token
func returnKeyVerification(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		err := respostas.MsgError("Método de assinatura inválido %v")
		return nil, fmt.Errorf(err.Error(), token.Header["alg"])
	}

	return config.SecretKey, nil
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)

	token, erro := jwt.Parse(tokenString, returnKeyVerification)

	if erro != nil {
		return 0, erro
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioId, nil
	}

	return 0, respostas.MsgError("Token inválido")
}
