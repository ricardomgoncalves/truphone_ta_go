package app

import (
	"github.com/jinzhu/gorm"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/postgres"
	"github.com/urfave/cli/v2"
	"log"
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
		postgresConnectionUrl, err := postgres.BuildConnectionString(
			opts.PostgresUrl,
			opts.PostgresSslMode,
			opts.PostgresDbName,
			opts.PostgresUser,
			opts.PostgresPassword,
		)
		if err != nil {
			return err
		}

		log.Println(postgresConnectionUrl)

		db, err := gorm.Open("postgres", postgresConnectionUrl)
		if err != nil {
			return err
		}

		if err := postgres.CreateTables(db); err != nil {
			return err
		}

		return nil
	}

	return app.Run(os.Args)
}
