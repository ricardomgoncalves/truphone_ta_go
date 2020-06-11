package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
)

type FamilyRepo interface {
	CreateFamily(ctx context.Context, fam family.Family) error
	GetFamilyById(ctx context.Context, id uuid.UUID) (*family.Family, error)
	UpdateFamilyById(ctx context.Context, id uuid.UUID, fam family.Family) error
	DeleteFamilyById(ctx context.Context, id uuid.UUID) error // RICARDO: REMOVE MEMBERS AS WELL
}

type MemberRepo interface {
	CreateMember(ctx context.Context, member family.Member) error
	GetMemberById(ctx context.Context, id uuid.UUID) (*family.Member, error)
	GetMembersByFamilyId(ctx context.Context, familyId uuid.UUID, offset *int, limit *int) ([]family.Member, error)
	UpdateMemberById(ctx context.Context, id uuid.UUID, member family.Member) error
	DeleteMemberById(ctx context.Context, id uuid.UUID) error
}
