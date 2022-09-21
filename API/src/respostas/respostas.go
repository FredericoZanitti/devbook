package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em JSON para requisicao
func JSON(w http.ResponseWriter, statuscode int, dados interface{}) {
	w.WriteHeader(statuscode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// Erro retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statuscode int, erro error) {
	JSON(w, statuscode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
