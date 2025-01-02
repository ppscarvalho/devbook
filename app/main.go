package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("Rodando webapp")
	utils.LoadTemplates()
	r := router.GenerateRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
