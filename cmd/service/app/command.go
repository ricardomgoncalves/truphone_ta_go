package app

import (
	"github.com/urfave/cli/v2"
	"os"
)

const (
	name = "truphone-family"
)

func New() error {
	app := cli.NewApp()
	app.Name = name

	opts := NewOptions()
	opts.Apply(app)

	app.Action = func(context *cli.Context) error {
		service := NewServiceApp()
		return service.Run(opts)
	}

	return app.Run(os.Args)
}
