package flags

import "github.com/urfave/cli/v2"

const (
	postgresUserName = "postgres_user"
	postgresUserShort = "puser"
	postgresUserEnvVarName = "POSTGRES_USER"
	postgresUserRequired = true
)

func PostgresUser(value *string) cli.Flag {
	return &cli.StringFlag{
		Name:        postgresUserName,
		Aliases:     []string{postgresUserShort},
		Usage:       "--"+postgresUserName+" user",
		EnvVars:     []string{postgresUserEnvVarName},
		Required:    postgresUserRequired,
		Destination: value,
	}
}
