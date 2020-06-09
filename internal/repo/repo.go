package repo

import "github.com/ricardomgoncalves/truphone_ta_go/pkg/family"

type FamilyRepo interface {
	GetFamilyById(id int) (family.Family, error)
}

type MemberRepo interface {
	GetMemberById(id int) (family.Member, error)
	GetMembersByFamilyId(familyId int) ([]family.Member, error)
}
