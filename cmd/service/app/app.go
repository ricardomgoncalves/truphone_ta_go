// Package classification Family API.
//
// the purpose of this application is to provide a management support
// of families.
//
//     Schemes: http, https
//     Host: 127.0.0.1:8080
//     BasePath: /truphone
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Ricardo Goncalves<ricardo.341928@hotmail.com> http://localhost
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package app

import (
	"github.com/ricardomgoncalves/truphone_ta_go/internal/postgres"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/routes"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/service"
	"log"
	"net/http"
	"time"
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

	db, err := postgres.TryConnectToDB(postgresConnectionUrl)
	if err != nil {
		return err
	}

	if err := postgres.CreateTables(db); err != nil {
		return err
	}

	if err := postgres.Populate(db); err != nil {
		return err
	}

	repo := postgres.NewPostgresRepo(db)
	svc, err := service.NewFamilyService(repo, repo)
	if err != nil {
		return err
	}

	router := routes.Router(svc)
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Println("***   starting server   ***")
	return srv.ListenAndServe()
}
