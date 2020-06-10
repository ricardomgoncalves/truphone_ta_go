package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"time"
)

type memberRow struct {
	*family.Member

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newMemberRow(member *family.Member) *memberRow {
	return &memberRow{Member: member}
}

func (m *memberRow) Value() *family.Member {
	return m.Member
}

func (*memberRow) TableName() string {
	return "members"
}

func (m *memberRow) BeforeSave(scope *gorm.Scope) error {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
	}
	m.UpdatedAt = time.Now()
	return nil
}

func (m *memberRow) AfterFind(scope *gorm.Scope) error {
	return nil
}
