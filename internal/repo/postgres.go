package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
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

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

var (
	_ FamilyRepo = (PostgresRepo)(nil)
	_ MemberRepo = (PostgresRepo)(nil)
)

func (p PostgresRepo) GetFamilyById(id int) (family.Family, error) {
	return family.Family{}, errors.New("not implemented")
}

func (p PostgresRepo) GetMemberById(id int) (family.Member, error) {
	return family.Member{}, errors.New("not implemented")
}

func (p PostgresRepo) GetMembersByFamilyId(familyId int) ([]family.Member, error) {
	return []family.Member{}, errors.New("not implemented")
}

func (p *PostgresRepo) checkFamilyError(err error) error {
	switch e2 := err.(type) {
	case nil:
		return nil
	case *pq.Error:
		switch e2.Code {
		case "40002":
			fallthrough
		case "42710":
			fallthrough
		case "23505":
			return errors.Wrap(err, family.ErrorFamilyAlreadyExists)
		case "55006":
			return errors.Wrap(err, family.ErrorFamilyLocked)
		}
		return errors.Wrap(err, family.ErrorFamilyUnknown)
	default:
		if err == gorm.ErrRecordNotFound {
			return errors.Wrap(err, family.ErrorFamilyNotFound)
		}
		return errors.Wrap(err, family.ErrorFamilyUnknown)
	}
}

func (p *PostgresRepo) checkMemberError(err error) error {
	switch e2 := err.(type) {
	case nil:
		return nil
	case *pq.Error:
		switch e2.Code {
		case "40002":
			fallthrough
		case "42710":
			fallthrough
		case "23505":
			return errors.Wrap(err, family.ErrorMemberAlreadyExists)
		case "55006":
			return errors.Wrap(err, family.ErrorMemberLocked)
		}
		return errors.Wrap(err, family.ErrorMemberUnknown)
	default:
		if err == gorm.ErrRecordNotFound {
			return errors.Wrap(err, family.ErrorMemberNotFound)
		}
		return errors.Wrap(err, family.ErrorMemberUnknown)
	}
}
