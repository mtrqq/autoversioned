package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "net/http"
    "os"

    lib "github.com/mtrqq/autoversioning/api/internal"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

var (
    secret = "none"
    port   = 8080
)

func getVersion(w http.ResponseWriter, req *http.Request) {
    log.Info().Msgf("Responding with version %s", lib.Version)

    response := map[string]string{
        "version": lib.Version,
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

    log.Info().Msgf("Starting http server at port %d", port)
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
