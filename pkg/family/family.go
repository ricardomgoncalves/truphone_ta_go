package family

import (
	"encoding/json"
	"errors"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/countrycode"
	"time"
)

var (
	ErrInvalidCountryCode = errors.New("invalid country code")
)

type Family struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
}

func (a *Family) UnmarshalJSON(b []byte) error {
	if !countrycode.IsValid(a.CountryCode) {
		return ErrInvalidCountryCode
	}

	if err := json.Unmarshal(b, a); err != nil {
		return err
	}
	return nil
}

func (a Family) MarshalJSON() ([]byte, error) {
	return json.Marshal(a)
}

type Member struct {
	Id         int       `json:"id"`
	FamilyId   int       `json:"family_id"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	FatherId   int       `json:"father_id"`
	MotherId   int       `json:"mother_id"`
	SpouseId   int       `json:"spouse_id"`
	Birthday   time.Time `json:"birthday"`
}

type memberRaw struct {
	Id         int    `json:"id"`
	FamilyId   int    `json:"family_id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	FatherId   int    `json:"father_id"`
	MotherId   int    `json:"mother_id"`
	SpouseId   int    `json:"spouse_id"`
	Birthday   string `json:"birthday"`
}

func (a Member) toRaw() memberRaw {
	return memberRaw{
		Id:         a.Id,
		FamilyId:   a.FamilyId,
		FirstName:  a.FirstName,
		MiddleName: a.MiddleName,
		LastName:   a.LastName,
		FatherId:   a.FatherId,
		MotherId:   a.MotherId,
		SpouseId:   a.SpouseId,
		Birthday:   a.Birthday.Format(time.RFC3339),
	}
}

func (a memberRaw) Parse() (Member, error) {
	birthday, err := time.Parse(time.RFC3339, a.Birthday)
	if err != nil {
		return Member{}, err
	}

	return Member{
		Id:         a.Id,
		FamilyId:   a.FamilyId,
		FirstName:  a.FirstName,
		MiddleName: a.MiddleName,
		LastName:   a.LastName,
		FatherId:   a.FatherId,
		MotherId:   a.MotherId,
		SpouseId:   a.SpouseId,
		Birthday:   birthday,
	}, nil
}

func (a *Member) UnmarshalJSON(b []byte) error {
	raw := memberRaw{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	member, err := raw.Parse()
	if err != nil {
		return err
	}

	*a = member
	return nil
}

func (a Member) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.toRaw())
}
