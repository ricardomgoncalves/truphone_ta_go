package flags

import "github.com/urfave/cli/v2"

const (
	postgresUrlName = "postgres_url"
	postgresUrlShort = "purl"
	postgresUrlEnvVarName = "POSTGRES_URL"
	postgresUrlRequired = true
)

func PostgresUrl(value *string) cli.Flag {
	return &cli.StringFlag{
		Name:        postgresUrlName,
		Aliases:     []string{postgresUrlShort},
		Usage:       "--"+postgresUrlName+" url",
		EnvVars:     []string{postgresUrlEnvVarName},
		Required:    postgresUrlRequired,
		Destination: value,
	}
}
