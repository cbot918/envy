package main

import (
	"os"

	"api/internal"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	HOST = "127.0.0.1:3501"
)

func main() {

	// init logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Str("foo", "bar").Msg("Hello world")

	app := internal.InitApi()

	if err := app.Listen(HOST); err != nil {
		panic(err)
	}
}
