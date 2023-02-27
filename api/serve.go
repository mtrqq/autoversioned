package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mtrqq/autoversioning/api/internal/config"
	"github.com/rs/zerolog/log"
)

func getVersion(w http.ResponseWriter, req *http.Request) {
	log.Info().Msgf("Responding with version %s", config.Version)

	response := map[string]string{
		"version": config.Version,
		"secret":  config.SecretValue,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getHealth(w http.ResponseWriter, req *http.Request) {
	response := map[string]string{
		"status": "ok",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getVersion)
	mux.HandleFunc("/health", getHealth)

	log.Info().Msgf("Starting http server at port %d", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), mux)
}
