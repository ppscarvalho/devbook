package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// ConfigurarCookie configura o cookie
func ConfigurarCookie() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Salvar(w http.ResponseWriter, ID, token string) error {
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}

	dadosCriptografados, erro := s.Encode("dados", dados)

	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCriptografados,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}

func ReadCookie(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}

	dados := make(map[string]string)
	if erro = s.Decode("dados", cookie.Value, &dados); erro != nil {
		return nil, erro
	}

	return dados, nil
}
