package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*func init() {
	key := make([]byte, 64)

	if _, erro := rand.Read(key); erro != nil {
		log.Fatal(erro)
	}

	keyBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(keyBase64)
}*/

func main() {
	config.Carregar()
	r := router.Gerar()
	fmt.Printf("Escutando na Porta: %d ", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
