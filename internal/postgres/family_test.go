package postgres

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/stretchr/testify/assert"
)

func TestNewFamilyRow(t *testing.T) {
	t.Run("should return a family row", func(t *testing.T) {
		fam := &family.Family{
			Id:          uuid.New(),
			Name:        "Family",
			CountryCode: "PT",
		}
		famRow := newFamilyRow(fam)
		assert.Equal(t, fam, famRow.Value())
	})
}

func TestFamilyRow_TableName(t *testing.T) {
	t.Run("should return string", func(t *testing.T) {
		famRow := newFamilyRow(nil)
		assert.Equal(t, "families", famRow.TableName())
	})
}

func TestFamilyRow_BeforeSave(t *testing.T) {
	t.Run("should set created and updated date the same", func(t *testing.T) {
		famRow := newFamilyRow(nil)
		err := famRow.BeforeSave(nil)
		assert.Nil(t, err)
		assert.False(t, famRow.CreatedAt.IsZero())
		assert.False(t, famRow.UpdatedAt.IsZero())
		assert.Equal(t, famRow.CreatedAt, famRow.UpdatedAt)
	})
	t.Run("should only set updated date", func(t *testing.T) {
		famRow := newFamilyRow(nil)
		famRow.CreatedAt = time.Now()
		err := famRow.BeforeSave(nil)
		assert.Nil(t, err)
		assert.False(t, famRow.CreatedAt.IsZero())
		assert.False(t, famRow.UpdatedAt.IsZero())
		assert.NotEqual(t, famRow.CreatedAt, famRow.UpdatedAt)
	})
}
