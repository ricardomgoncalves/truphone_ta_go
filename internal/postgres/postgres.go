package postgres

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/ricardomgoncalves/truphone_ta_go/internal/repo"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
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

func (p Repo) CreateFamily(ctx context.Context, fam family.Family) error {
	db := p.db.Set("ctx", ctx)

	row := newFamilyRow(&fam)

	if err := db.Create(row).Error; err != nil {
		return err
	}

	return nil
}

func (p Repo) GetFamilyById(ctx context.Context, id string) (*family.Family, error) {
	db := p.db.Set("ctx", ctx)

	fml := family.NewFamilyWithId(id)
	row := newFamilyRow(&fml)

	if err := db.Where("id = ?", id).First(row).Error; err != nil {
		return nil, p.checkFamilyError(err)
	}

	return row.Value(), nil
}

func (p Repo) ListFamilies(ctx context.Context, options ...repo.FilterOption) ([]family.Family, error) {
	db := p.db.Set("ctx", ctx)

	opts := repo.FilterOptions{}
	for _, option := range options {
		option(&opts)
	}

	if val := repo.GetLimit(opts); val != nil && *val != 0 {
		db = db.Limit(*val)
	}

	if val := repo.GetOffset(opts); val != nil && *val != 0 {
		db = db.Offset(*val)
	}

	if val := repo.GetCountryCode(opts); val != nil {
		db = db.Where("country_code = '" + *val + "'")
	}

	rows := make([]*familyRow, 0, 200)
	if err := db.Find(&rows).Error; err != nil {
		return nil, p.checkFamilyError(err)
	}

	results := make([]family.Family, len(rows))
	for i, row := range rows {
		if row == nil {
			continue
		}

		results[i] = *row.Value()
	}
	return results, nil
}

func (p Repo) UpdateFamilyById(ctx context.Context, id string, fam family.Family) error {
	famRow := newFamilyRow(&fam)

	db := p.db.
		Set("ctx", ctx).
		Model(famRow).
		Where("id = ?", id).
		Update(famRow)

	if db.Error != nil {
		return p.checkFamilyError(db.Error)
	}

	if db.RowsAffected == 0 {
		return p.checkFamilyError(family.ErrorFamilyNotFound)
	}

	return nil
}

func (p Repo) DeleteFamilyById(ctx context.Context, id string) error {
	fml := family.NewFamilyWithId(id)
	row := newFamilyRow(&fml)

	db := p.db.
		Set("ctx", ctx).
		Delete(row, "id = '"+id+"'")

	if db.Error != nil {
		return p.checkFamilyError(db.Error)
	}

	if db.RowsAffected == 0 {
		return p.checkFamilyError(family.ErrorFamilyNotFound)
	}

	return nil
}

func (p Repo) CreateMember(ctx context.Context, member family.Member) error {
	db := p.db.Set("ctx", ctx)

	row := newMemberRow(&member)
	if err := db.Create(row).Error; err != nil {
		return err
	}

	return nil
}

func (p Repo) GetMemberById(ctx context.Context, id string) (*family.Member, error) {
	db := p.db.Set("ctx", ctx)

	member := family.NewMemberWithId(id)
	row := newMemberRow(&member)

	if err := db.Where("id = ?", id).First(row).Error; err != nil {
		return nil, p.checkFamilyError(err)
	}

	return row.Value(), nil
}

func (p Repo) GetMembersByFamilyId(ctx context.Context, familyId string, offset *int, limit *int) ([]family.Member, error) {
	db := p.db.Set("ctx", ctx)

	if limit != nil && *limit != 0 {
		db = db.Limit(*limit)
	}

	if offset != nil && *offset != 0 {
		db = db.Offset(*offset)
	}

	db = db.Where("family_id = ?", familyId)

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

func (p Repo) UpdateMemberById(ctx context.Context, id string, member family.Member) error {
	memberRow := newMemberRow(&member)

	db := p.db.
		Set("ctx", ctx).
		Model(memberRow).
		Where("id = ?", id).
		Update(memberRow)

	if db.Error != nil {
		return p.checkMemberError(db.Error)
	}

	if db.RowsAffected == 0 {
		return p.checkMemberError(family.ErrorMemberNotFound)
	}

	return nil
}

func (p Repo) DeleteMemberById(ctx context.Context, id string) error {
	member := family.NewMemberWithId(id)
	row := newMemberRow(&member)

	db := p.db.
		Set("ctx", ctx).
		Delete(row, "id = '"+id+"'")

	if db.Error != nil {
		return p.checkMemberError(db.Error)
	}

	if db.RowsAffected == 0 {
		return p.checkMemberError(family.ErrorMemberNotFound)
	}

	return nil
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
		switch err {
		case family.ErrorFamilyNotFound:
			return err
		case gorm.ErrRecordNotFound:
			return errors.Wrap(err, family.ErrorFamilyNotFound)
		default:
			return errors.Wrap(err, family.ErrorFamilyUnknown)
		}
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
		switch err {
		case family.ErrorMemberNotFound:
			return err
		case gorm.ErrRecordNotFound:
			return errors.Wrap(err, family.ErrorMemberNotFound)
		default:
			return errors.Wrap(err, family.ErrorMemberUnknown)
		}
	}
}
