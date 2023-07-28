package internal

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	CLOUD_HOST string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Error().Msg("load .env failed")
	}
	return &Config{
		CLOUD_HOST: os.Getenv("CLOUD_HOST"),
	}
}
