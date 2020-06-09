package repo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuildConnectionString(t *testing.T) {
	t.Run("should return connection with port", func(t *testing.T) {
		input := []struct {
			host     string
			sslmode  string
			dbname   string
			user     string
			password string
			output   string
		}{
			{
				host:     "//localhost:5432",
				sslmode:  "disable",
				dbname:   "test",
				user:     "user",
				password: "pass",
				output:   "host=localhost port=5432 sslmode=disable dbname=test user=user password=pass",
			},
			{
				host:     "//localhost:2344",
				sslmode:  "disable",
				dbname:   "test",
				user:     "user",
				password: "pass",
				output:   "host=localhost port=2344 sslmode=disable dbname=test user=user password=pass",
			},
		}

		for _, in := range input {
			output, err := BuildConnectionString(in.host, in.sslmode, in.dbname, in.user, in.password)
			require.Nil(t, err)
			assert.Equal(t, in.output, output)
		}
	})
	t.Run("should return connection without port", func(t *testing.T) {
		input := []struct {
			host     string
			sslmode  string
			dbname   string
			user     string
			password string
			output   string
		}{
			{
				host:     "//localhost",
				sslmode:  "disable",
				dbname:   "test",
				user:     "user",
				password: "pass",
				output:   "host=localhost sslmode=disable dbname=test user=user password=pass",
			},
		}

		for _, in := range input {
			output, err := BuildConnectionString(in.host, in.sslmode, in.dbname, in.user, in.password)
			require.Nil(t, err)
			assert.Equal(t, in.output, output)
		}
	})
	t.Run("should return error", func(t *testing.T) {
		input := []struct {
			host     string
			sslmode  string
			dbname   string
			user     string
			password string
		}{
			{
				host:     ":foo",
				sslmode:  "disable",
				dbname:   "test",
				user:     "user",
				password: "pass",
			},
		}

		for _, in := range input {
			_, err := BuildConnectionString(in.host, in.sslmode, in.dbname, in.user, in.password)
			require.NotNil(t, err)
		}
	})
}
