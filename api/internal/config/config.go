package config

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	Version     = "2.4.2"
	SecretValue = "none"
	Port        = 8080
)

func init() {
	flag.IntVar(&Port, "port", 8080, "port to listen at")
	flag.Parse()

	if secretEnv, exists := os.LookupEnv("SECRET_VALUE"); exists {
		SecretValue = secretEnv
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}
