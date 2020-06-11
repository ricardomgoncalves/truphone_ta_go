package postgres

import (
	"testing"
	"time"

	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/stretchr/testify/assert"
)

func TestNewMemberRow(t *testing.T) {
	t.Run("should return a family row", func(t *testing.T) {
		fam := &family.Member{
			Id: "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da038",
		}
		famRow := newMemberRow(fam)
		assert.Equal(t, fam, famRow.Value())
	})
}

func TestMemberRow_TableName(t *testing.T) {
	t.Run("should return string", func(t *testing.T) {
		famRow := newMemberRow(nil)
		assert.Equal(t, "members", famRow.TableName())
	})
}

func TestMemberRow_BeforeSave(t *testing.T) {
	t.Run("should set created and updated date the same", func(t *testing.T) {
		famRow := newMemberRow(nil)
		err := famRow.BeforeSave(nil)
		assert.Nil(t, err)
		assert.False(t, famRow.CreatedAt.IsZero())
		assert.False(t, famRow.UpdatedAt.IsZero())
		assert.Equal(t, famRow.CreatedAt, famRow.UpdatedAt)
	})
	t.Run("should only set updated date", func(t *testing.T) {
		famRow := newMemberRow(nil)
		famRow.CreatedAt = time.Now()
		err := famRow.BeforeSave(nil)
		assert.Nil(t, err)
		assert.False(t, famRow.CreatedAt.IsZero())
		assert.False(t, famRow.UpdatedAt.IsZero())
		assert.NotEqual(t, famRow.CreatedAt, famRow.UpdatedAt)
	})
}
