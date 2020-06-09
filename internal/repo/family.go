package repo

import (
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"time"
)

type familyRow struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	CountryCode string    `json:"country_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func fromFamily(family family.Family) familyRow {
	return familyRow{
		Id:          family.Id,
		Name:        family.Name,
		CountryCode: family.CountryCode,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
