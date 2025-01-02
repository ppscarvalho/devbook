package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

/*
func init() {
	kashkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(kashkey)

	blockkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(blockkey)
}*/

func main() {
	config.Carregar()
	cookies.ConfigurarCookie()
	utils.LoadTemplates()
	r := router.GenerateRouter()

	fmt.Printf("Aplicação executando na porta: %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
