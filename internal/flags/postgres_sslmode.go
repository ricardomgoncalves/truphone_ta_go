package flags

import "github.com/urfave/cli/v2"

const (
	postgresDbNameName = "postgres_dbname"
	postgresDbNameShort = "pdb"
	postgresDbNameEnvVarName = "POSTGRES_DBNAME"
	postgresDbNameRequired = true
)

func PostgresDbName(value *string) cli.Flag {
	return &cli.StringFlag{
		Name:        postgresDbNameName,
		Aliases:     []string{postgresDbNameShort},
		Usage:       "--"+postgresDbNameName+" name",
		EnvVars:     []string{postgresDbNameEnvVarName},
		Required:    postgresDbNameRequired,
		Destination: value,
	}
}
