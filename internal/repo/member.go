package repo

import "time"

type memberRaw struct {
	Id         int       `json:"id"`
	FamilyId   int       `json:"family_id"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	FatherId   int       `json:"father_id"`
	MotherId   int       `json:"mother_id"`
	SpouseId   int       `json:"spouse_id"`
	Birthday   time.Time `json:"birthday"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
