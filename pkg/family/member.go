package family

import (
	"encoding/json"
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
	// max length: 30
	FirstName string `json:"first_name"`

	// the middle name for this member
	//
	// required: true
	// min length: 3
	// max length: 30
	MiddleName string `json:"middle_name,omitempty"`

	// the last name for this member
	//
	// required: true
	// min length: 3
	// max length: 30
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

func (m Member) toRaw() memberRaw {
	return memberRaw{
		Id:         m.Id,
		FamilyId:   m.FamilyId,
		FirstName:  m.FirstName,
		MiddleName: m.MiddleName,
		LastName:   m.LastName,
		FatherId:   m.FatherId,
		MotherId:   m.MotherId,
		SpouseId:   m.SpouseId,
		Birthday:   m.Birthday.Format(time.RFC3339),
	}
}

func (m memberRaw) parse() Member {
	member := Member{
		Id:         m.Id,
		FamilyId:   m.FamilyId,
		FirstName:  m.FirstName,
		MiddleName: m.MiddleName,
		LastName:   m.LastName,
		FatherId:   m.FatherId,
		MotherId:   m.MotherId,
		SpouseId:   m.SpouseId,
		Birthday:   time.Time{},
	}

	birthday, err := time.Parse(time.RFC3339, m.Birthday)
	if err == nil {
		member.Birthday = birthday
	}

	return member
}

func (m *Member) UnmarshalJSON(b []byte) error {
	raw := memberRaw{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	*m = raw.parse()
	return nil
}

func (m *Member) Patch(member Member) {
	if m == nil {
		return
	}

	if member.FamilyId != "" {
		m.FamilyId = member.FamilyId
	}

	if member.FirstName != "" {
		m.FirstName = member.FirstName
	}

	if member.MiddleName != "" {
		m.MiddleName = member.MiddleName
	}

	if member.LastName != "" {
		m.LastName = member.LastName
	}

	if member.FatherId != nil {
		val := *member.FatherId
		m.FatherId = &val
	}

	if member.MotherId != nil {
		val := *member.MotherId
		m.MotherId = &val
	}

	if member.SpouseId != nil {
		val := *member.SpouseId
		m.SpouseId = &val
	}

	if !member.Birthday.IsZero() {
		m.Birthday = member.Birthday
	}
}

func (m Member) HasCommonName(member Member) bool {
	if m.FirstName == member.FirstName {
		return true
	}

	if m.MiddleName == member.MiddleName {
		return true
	}

	if m.LastName == member.LastName {
		return true
	}

	return false
}

func (m Member) HasSimilarBirthday(member Member) bool {
	if m.Birthday == member.Birthday {
		return true
	}

	return false
}

func (m Member) IsMissingMother() bool {
	return m.MotherId == nil
}

func (m Member) IsMissingFather() bool {
	return m.FatherId == nil
}

func (m Member) HasSameMother(member Member) bool {
	if m.MotherId == member.MotherId {
		return true
	}

	if m.MotherId == nil || member.MotherId == nil {
		return false
	}

	if *m.MotherId == *member.MotherId {
		return true
	}

	return false
}

func (m Member) HasSameFather(member Member) bool {
	if m.FatherId == member.FatherId {
		return true
	}

	if m.FatherId == nil || member.FatherId == nil {
		return false
	}

	if *m.FatherId == *member.FatherId {
		return true
	}

	return false
}