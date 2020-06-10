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

func (m *familyRow) BeforeSave(scope *gorm.Scope) error {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
	}
	m.UpdatedAt = time.Now()
	return nil
}

func (m *familyRow) AfterFind(scope *gorm.Scope) error {
	return nil
}
