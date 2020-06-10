package postgres

import (
	"net/url"
	"strings"
)

func BuildConnectionString(host, sslMode, dbName, user, password string) (string, error) {
	return buildConnStr(host, sslMode, dbName, user, password)
}

func buildConnStr(host string, ssl, store, user, password string) (string, error) {
	hostUrl, err := url.Parse(host)
	if err != nil {
		return "", err
	}

	connStr := new(strings.Builder)
	connStr.WriteString("host=")
	connStr.WriteString(strings.TrimSpace(hostUrl.Hostname()))

	if port := hostUrl.Port(); port != "" {
		connStr.WriteString(" port=")
		connStr.WriteString(strings.TrimSpace(port))
	}
	connStr.WriteString(" sslmode=")
	connStr.WriteString(strings.TrimSpace(ssl))

	connStr.WriteString(" dbname=")
	connStr.WriteString(strings.TrimSpace(store))

	connStr.WriteString(" user=")
	connStr.WriteString(strings.TrimSpace(user))

	if password != "" {
		connStr.WriteString(" password=")
		connStr.WriteString(strings.TrimSpace(password))
	}

	return connStr.String(), nil
}
