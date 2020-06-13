package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"log"
	"net/url"
	"strings"
	"time"
)

func TryConnectToDB(postgresConnectionUrl string) (*gorm.DB, error) {
	tries := 0
	for tries < 5 {
		db, err := gorm.Open("postgres", postgresConnectionUrl)
		if err == nil {
			log.Println("connected to database")
			return db, nil
		}
		tries++
		log.Println("failed connecting to postgres:", err.Error())
		log.Println("retrying... (try:", tries, ")")
		time.Sleep(time.Second * 5)
	}

	return nil, errors.New("maximum tries hit")
}

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
