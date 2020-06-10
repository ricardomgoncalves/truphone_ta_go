package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/repo"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
)

type Repo struct {
	db *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

var (
	_ repo.FamilyRepo = (*Repo)(nil)
	_ repo.MemberRepo = (*Repo)(nil)
)

func (p Repo) GetFamilyById(ctx context.Context, id uuid.UUID) (*family.Family, error) {
	db := p.db.Set("ctx", ctx)

	fml := family.NewFamilyWithId(id)
	row := newFamilyRow(&fml)

	if err := db.Where("id = ?", id.String()).First(row).Error; err != nil {
		return nil, p.checkFamilyError(err)
	}

	return row.Value(), nil
}

func (p Repo) GetMemberById(ctx context.Context, id uuid.UUID) (*family.Member, error) {
	db := p.db.Set("ctx", ctx)

	member := family.NewMemberWithId(id)
	row := newMemberRow(&member)

	if err := db.Where("id = ?", id.String()).First(row).Error; err != nil {
		return nil, p.checkFamilyError(err)
	}

	return row.Value(), nil
}

func (p Repo) GetMembersByFamilyId(ctx context.Context, familyId uuid.UUID, offset *int, limit *int) ([]family.Member, error) {
	db := p.db.Set("ctx", ctx)

	if limit != nil && *limit != 0 {
		db = db.Limit(*limit)
	}

	if offset != nil && *offset != 0 {
		db = db.Offset(*offset)
	}

	db = db.Where("family_id = ?", familyId.String())

	rows := make([]*memberRow, 0, 200)
	if err := db.Find(&rows).Error; err != nil {
		return nil, p.checkMemberError(err)
	}

	results := make([]family.Member, len(rows))
	for i, row := range rows {
		if row == nil {
			continue
		}

		results[i] = *row.Value()
	}

	return results, nil
}

func (p *Repo) checkFamilyError(err error) error {
	switch e2 := err.(type) {
	case nil:
		return nil
	case *pq.Error:
		switch e2.Code {
		case "40002":
			fallthrough
		case "42710":
			fallthrough
		case "23505":
			return errors.Wrap(err, family.ErrorFamilyAlreadyExists)
		case "55006":
			return errors.Wrap(err, family.ErrorFamilyLocked)
		}
		return errors.Wrap(err, family.ErrorFamilyUnknown)
	default:
		if err == gorm.ErrRecordNotFound {
			return errors.Wrap(err, family.ErrorFamilyNotFound)
		}
		return errors.Wrap(err, family.ErrorFamilyUnknown)
	}
}

func (p *Repo) checkMemberError(err error) error {
	switch e2 := err.(type) {
	case nil:
		return nil
	case *pq.Error:
		switch e2.Code {
		case "40002":
			fallthrough
		case "42710":
			fallthrough
		case "23505":
			return errors.Wrap(err, family.ErrorMemberAlreadyExists)
		case "55006":
			return errors.Wrap(err, family.ErrorMemberLocked)
		}
		return errors.Wrap(err, family.ErrorMemberUnknown)
	default:
		if err == gorm.ErrRecordNotFound {
			return errors.Wrap(err, family.ErrorMemberNotFound)
		}
		return errors.Wrap(err, family.ErrorMemberUnknown)
	}
}
