package app

import (
	"os"

	"github.com/alecthomas/kong"
	"go.uber.org/fx"
)

type CLI struct {
	fx.Out

	Config string `optional:"" short:"c" type:"existingfile" placeholder:"PATH" help:"Path to the config file." name:"config_path"`
}

func NewCLI() (cli CLI, err error) {
	k, err := kong.New(&cli)
	if err != nil {
		return CLI{}, err
	}

	_, err = k.Parse(os.Args[1:])
	if err != nil {
		return CLI{}, err
	}

	return cli, nil
}
