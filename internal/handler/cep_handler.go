package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/eucaiot/cep-api/internal/model"
)

const (
	VIA_CEP_URL = "https://viacep.com.br/ws/%s/json/"
)

func BuscaCep(w http.ResponseWriter, r *http.Request) {

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Param cep não informado"))
		return
	}

	res, err := callViaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Um erro inesperado ocorreu"))
		slog.Error("[ViaCep] - erro ao chamar API.", "Erro", err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Um erro inesperado ocorreu"))
		slog.Error("[CEP-API] - Erro ao gerar JSON", "Erro", err.Error())
		return
	}

}

func callViaCep(cep string) (*model.Cep, error){
	res, err := http.Get(fmt.Sprintf(VIA_CEP_URL, cep))
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	var cepDto model.Cep
	err = json.Unmarshal(data, &cepDto)
	if err != nil {
		return nil, err
	}

	return &cepDto, nil
}