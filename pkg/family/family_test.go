package family

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFamilyWithId(t *testing.T) {
	t.Run("should return family with id", func(t *testing.T) {
		id := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da038"
		family := NewFamilyWithId(id)
		assert.Equal(t, id, family.Id)
	})
}

func TestFamily_Patch(t *testing.T) {
	t.Run("should not do anything on nil family", func(t *testing.T) {
		var fam *Family
		fam.Patch(Family{})
	})
	t.Run("should only update name", func(t *testing.T) {
		fam := &Family{
			Id:          "id",
			Name:        "Family 1",
			CountryCode: "PT",
		}
		fam.Patch(Family{Name: "updated"})
		assert.Equal(t, "id", fam.Id)
		assert.Equal(t, "updated", fam.Name)
		assert.Equal(t, "PT", fam.CountryCode)
	})
	t.Run("should only updated country code", func(t *testing.T) {
		fam := &Family{
			Id:          "id",
			Name:        "Family 1",
			CountryCode: "PT",
		}
		fam.Patch(Family{CountryCode: "ES"})
		assert.Equal(t, "id", fam.Id)
		assert.Equal(t, "Family 1", fam.Name)
		assert.Equal(t, "ES", fam.CountryCode)
	})
	t.Run("should only updated name and country code", func(t *testing.T) {
		fam := &Family{
			Id:          "id",
			Name:        "Family 1",
			CountryCode: "PT",
		}
		fam.Patch(Family{
			Id:          "should_not_update",
			Name:        "updated",
			CountryCode: "ES",
		})
		assert.Equal(t, "id", fam.Id)
		assert.Equal(t, "updated", fam.Name)
		assert.Equal(t, "ES", fam.CountryCode)
	})
}
