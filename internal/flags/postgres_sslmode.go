package flags

import "github.com/urfave/cli/v2"

const (
	postgresSslModeName = "postgres_sslmode"
	postgresSslModeShort = "pssl"
	postgresSslModeEnvVarName = "POSTGRES_SSLMODE"
	postgresSslModeRequired = true
)

func PostgresSslMode(value *string) cli.Flag {
	return &cli.StringFlag{
		Name:        postgresSslModeName,
		Aliases:     []string{postgresSslModeShort},
		Usage:       "--"+postgresSslModeName+" mode",
		EnvVars:     []string{postgresSslModeEnvVarName},
		Required:    postgresSslModeRequired,
		Destination: value,
	}
}
