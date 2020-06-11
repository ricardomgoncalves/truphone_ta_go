package family

import (
	"encoding/json"
	"log"
	"time"
)

// Member
//
// A member can have a father, mother and spouse.
//
// swagger:model
type Member struct {
	// the id for this member
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	Id string `json:"id"`

	// the family id for this member
	//
	// required: true
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	FamilyId string `json:"family_id,omitempty"`

	// the first name for this member
	//
	// required: true
	// min length: 3
	FirstName string `json:"first_name"`

	// the middle name for this member
	//
	// required: true
	// min length: 3
	MiddleName string `json:"middle_name,omitempty"`

	// the last name for this member
	//
	// required: true
	// min length: 3
	LastName string `json:"last_name"`

	// the father id for this member
	//
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	FatherId *string `json:"father_id,omitempty"`

	// the mother id for this member
	//
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	MotherId *string `json:"mother_id,omitempty"`

	// the spouse id for this member
	//
	// example: 8957bf28-aea0-47de-abe0-d4c5ea593ec6
	SpouseId *string `json:"spouse_id,omitempty"`

	// the birthday for this member
	//
	// required: true
	// example: 2012-02-03T00:04:05Z
	Birthday time.Time `json:"birthday"`
}

func NewMemberWithId(id string) Member {
	return Member{Id: id}
}

type memberRaw struct {
	Id         string  `json:"id"`
	FamilyId   string  `json:"family_id"`
	FirstName  string  `json:"first_name"`
	MiddleName string  `json:"middle_name"`
	LastName   string  `json:"last_name"`
	FatherId   *string `json:"father_id"`
	MotherId   *string `json:"mother_id"`
	SpouseId   *string `json:"spouse_id"`
	Birthday   string  `json:"birthday"`
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

func (a memberRaw) parse() (Member, error) {
	log.Println(a)
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

	member, err := raw.parse()
	if err != nil {
		return err
	}

	*a = member
	return nil
}
