package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)
//init gera a chave secret, que é usada uma única vez e salva no .env
//func init() {
//	chave := make([]byte, 64)
//
//	if _, erro := rand.Read(chave); erro != nil {
//		log.Fatal(erro)
//	}
//
//	stringBase64 := base64.StdEncoding.EncodeToString(chave)
//	fmt.Println(stringBase64)
//}

func main() {
	config.Carregar()
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",config.Porta), r))
}