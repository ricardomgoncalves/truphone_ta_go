package family

// Family
//
// A family can have a name and a country.
//
// swagger:model
type Family struct {
	// the id for this family
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the name for this family
	//
	// required: true
	// min length: 3
	// max length: 30
	Name string `json:"name"`

	// the country code for this family
	//
	// required: true
	// max length: 2
	// min length: 2
	CountryCode string `json:"country_code"`
}

func NewFamilyWithId(id string) Family {
	return Family{Id: id}
}

func (f *Family) Patch(fam Family) {
	if f == nil {
		return
	}

	if fam.Name != "" {
		f.Name = fam.Name
	}

	if fam.CountryCode != "" {
		f.CountryCode = fam.CountryCode
	}
}
