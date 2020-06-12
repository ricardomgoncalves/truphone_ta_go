package repo

import (
	"context"

	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
)

type FamilyRepo interface {
	CreateFamily(ctx context.Context, fam family.Family) error
	GetFamilyById(ctx context.Context, id string) (*family.Family, error)
	ListFamilies(ctx context.Context, opts ...FilterOption) ([]family.Family, error)
	UpdateFamilyById(ctx context.Context, id string, fam family.Family) error
	DeleteFamilyById(ctx context.Context, id string) error
}

type MemberRepo interface {
	CreateMember(ctx context.Context, member family.Member) error
	GetMemberById(ctx context.Context, id string) (*family.Member, error)
	ListMembers(ctx context.Context, options ...FilterOption) ([]family.Member, error)
	UpdateMemberById(ctx context.Context, id string, member family.Member) error
	DeleteMemberById(ctx context.Context, id string) error
}
