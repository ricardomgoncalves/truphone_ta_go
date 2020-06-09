package flags

import "github.com/urfave/cli/v2"

const (
	postgresPasswordName = "postgres_password"
	postgresPasswordShort = "ppw"
	postgresPasswordEnvVarName = "POSTGRES_PASSWORD"
	postgresPasswordRequired = true
)

func PostgresPassword(value *string) cli.Flag {
	return &cli.StringFlag{
		Name:        postgresPasswordName,
		Aliases:     []string{postgresPasswordShort},
		Usage:       "--"+postgresPasswordName+" password",
		EnvVars:     []string{postgresPasswordEnvVarName},
		Required:    postgresPasswordRequired,
		Destination: value,
	}
}
