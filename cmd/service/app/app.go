package app

import (
	"github.com/ricardomgoncalves/truphone_ta_go/internal/postgres"
	"log"
)

type ServiceApp struct{}

func NewServiceApp() ServiceApp {
	return ServiceApp{}
}

func (ServiceApp) Run(opts Options) error {
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
	return nil
}
