package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
)

type FamilyRepo interface {
	GetFamilyById(ctx context.Context, id uuid.UUID) (*family.Family, error)
}

type MemberRepo interface {
	GetMemberById(ctx context.Context, id uuid.UUID) (*family.Member, error)
	GetMembersByFamilyId(ctx context.Context, familyId uuid.UUID, offset *int, limit *int) ([]family.Member, error)
}
