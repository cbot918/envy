package main

import (
	"cloudapi/internal"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var lg = fmt.Println
var lgj = func(a any) {
	fmt.Printf("%+v\n", a)
}

func main() {
	// init config
	cfg := internal.NewConfig()

	// init logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// init http server
	api, err := internal.NewCloudAPI(cfg, log.Logger)
	if err != nil {
		panic(err)
	}
	if err := api.Server.Listen(cfg.HOST); err != nil {
		panic(err)
	}
}
