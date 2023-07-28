package internal

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

type Cli struct {
	Cfg *Config
}

func NewCli(cfg *Config) *Cli {
	return &Cli{
		Cfg: cfg,
	}
}

func (c *Cli) Run() error {
	args := os.Args

	if len(args) == 1 {
		return fmt.Errorf("need one command")
	}
	if len(args) == 3 {
		return fmt.Errorf("cmds out of range")
	}

	switch args[1] {
	case INSTALL:
		log.Info().Msg("cmd: eny install [package]")
		log.Info().Msg("ex: eny install redis")
	case LOGIN:
		log.Info().Msg("cmd: eny login [id] [token]")
		log.Info().Msg("ex: eny login yale918 12345")
		if args[2] != "" && args[3] != "" {

			client := resty.New()
			resp, err := client.R().Get(c.Cfg.CLOUD_HOST + "/login")
			if err != nil {
				return err
			}
			Lgj(resp)

		}
	default:
		return fmt.Errorf("cmd not recognize")
	}

	return nil
}
