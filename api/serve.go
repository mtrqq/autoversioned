package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	version = "2.0.1"
	secret  = "none"
	port    = 8080
)

func getVersion(w http.ResponseWriter, req *http.Request) {
	log.Info().Msgf("Responding with version %s", version)

	response := map[string]string{
		"version": version,
		"secret":  secret,
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

	log.Info().Msg("Starting http server at port 8080")
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func init() {
	flag.IntVar(&port, "port", 8080, "port to listen at")
	flag.Parse()

	if secretEnv, exists := os.LookupEnv("SECRET_VALUE"); exists {
		secret = secretEnv
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
