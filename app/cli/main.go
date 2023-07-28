package main

import (
	"cli/internal"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var lg = fmt.Println

func main() {
	// init log
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// init config
	cfg := internal.NewConfig()

	cli := internal.NewCli(cfg)
	if err := cli.Run(); err != nil {
		log.Error().Msg(err.Error())
	}
}
