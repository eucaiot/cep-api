package main

import (
	"log/slog"
	"net/http"
	"github.com/eucaiot/cep-api/internal/handler"
)

func main() {

	http.HandleFunc("/", handler.Start)
	http.HandleFunc("/cep", handler.BuscaCep)
	slog.Info("started...")
	http.ListenAndServe(":8080", nil)
}
