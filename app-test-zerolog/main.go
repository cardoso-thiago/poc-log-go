package main

import (
	"encoding/json"
	"net/http"

	"github.com/cardoso-thiago/quicklog"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	quicklog.GetLogger().Info().
		Str("cpf", "333.333.333-33").
		Msg("Teste com máscara de CPF")

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(Response{Message: "Hello World!"})
}

func main() {
	http.HandleFunc("/", helloHandler)
	//Adiciona um log de informação com um novo atributo `service`
	quicklog.GetLogger().Info().
		Str("service", "poc-log-go").
		Msg("Iniciando o servidor na porta 8888")

	if err := http.ListenAndServe(":8888", nil); err != nil {
		//Adiciona um log de erro com o objeto de erro e uma mensagem customizada
		quicklog.GetLogger().
			Err(err).
			Msgf("Não foi possível iniciar o servidor: %s", err.Error())
	}
}
