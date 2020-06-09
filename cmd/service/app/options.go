package app

import "github.com/urfave/cli/v2"

type Options struct {

}

func NewOptions() Options {
	return Options{}
}

func (opts Options) Apply(app *cli.App) {
	newFlags := []cli.Flag {

	}

	app.Flags = append(app.Flags, newFlags...)
}
