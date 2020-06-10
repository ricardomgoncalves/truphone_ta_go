package family

import (
	"github.com/google/uuid"
)

type Family struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	CountryCode string    `json:"country_code"`
}

func NewFamilyWithId(id uuid.UUID) Family {
	return Family{Id: id}
}
