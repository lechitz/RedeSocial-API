package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON retorna uma resposta para a requisição. Essa função tem 3 parâmetros: http.ResponseWrite que escreve as respostas na saída, statusCode, dados de interface
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatalln(erro)
	}
}

//Erro retorna um erro em formato JSON.
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
