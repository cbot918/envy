package internal

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	HOST string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Error().Msg("load .env failed")
	}
	return &Config{
		HOST: os.Getenv("HOST"),
	}
}
