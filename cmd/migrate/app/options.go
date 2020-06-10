package app

import (
	"github.com/ricardomgoncalves/truphone_ta_go/internal/flags"
	"github.com/urfave/cli/v2"
)

type Options struct {
	PostgresUrl      string
	PostgresSslMode  string
	PostgresDbName   string
	PostgresUser     string
	PostgresPassword string
}

func NewOptions() *Options {
	return &Options{}
}

func (opts *Options) Apply(app *cli.App) {
	newFlags := []cli.Flag{
		flags.PostgresUrl(&opts.PostgresUrl),
		flags.PostgresSslMode(&opts.PostgresSslMode),
		flags.PostgresDbName(&opts.PostgresDbName),
		flags.PostgresUser(&opts.PostgresUser),
		flags.PostgresPassword(&opts.PostgresPassword),
	}

	app.Flags = append(app.Flags, newFlags...)
}
