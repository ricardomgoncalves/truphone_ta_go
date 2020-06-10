package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"time"
)

type familyRow struct {
	*family.Family

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newFamilyRow(fml *family.Family) *familyRow {
	return &familyRow{Family: fml}
}

func (m *familyRow) Value() *family.Family {
	return m.Family
}

func (*familyRow) TableName() string {
	return "families"
}

func (m *familyRow) BeforeSave(_ *gorm.Scope) error {
	m.UpdatedAt = time.Now()

	if m.CreatedAt.IsZero() {
		m.CreatedAt = m.UpdatedAt
	}

	return nil
}
