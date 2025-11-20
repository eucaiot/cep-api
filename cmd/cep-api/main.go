package main

import (
	"log/slog"
	"net/http"
)

func main() {

	defer slog.Info("started...")
	http.ListenAndServe(":8080", nil)
}
