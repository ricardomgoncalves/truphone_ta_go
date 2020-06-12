package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithOffset(t *testing.T) {
	t.Run("should set offset", func(t *testing.T) {
		opts := FilterOptions{}
		WithOffset(1)(&opts)
		assert.Equal(t, uint32(1), *opts.offset)
	})
}

func TestGetOffset(t *testing.T) {
	t.Run("should get limit", func(t *testing.T) {
		offset := uint32(1)
		opts := FilterOptions{
			offset: &offset,
		}
		assert.Equal(t, uint32(1), *GetOffset(opts))
	})
}

func TestWithLimit(t *testing.T) {
	t.Run("should set limit", func(t *testing.T) {
		opts := FilterOptions{}
		WithLimit(1)(&opts)
		assert.Equal(t, uint32(1), *opts.limit)
	})
}

func TestGetLimit(t *testing.T) {
	t.Run("should get limit", func(t *testing.T) {
		limit := uint32(1)
		opts := FilterOptions{
			limit: &limit,
		}
		assert.Equal(t, uint32(1), *GetLimit(opts))
	})
}

func TestWithCountryCode(t *testing.T) {
	t.Run("should set country", func(t *testing.T) {
		opts := FilterOptions{}
		WithCountryCode("PT")(&opts)
		assert.Equal(t, "PT", *opts.countryCode)
	})
}

func TestGetCountryCode(t *testing.T) {
	t.Run("should get country", func(t *testing.T) {
		country := "PT"
		opts := FilterOptions{
			countryCode: &country,
		}
		assert.Equal(t, "PT", *GetCountryCode(opts))
	})
}

func TestWithFamilyId(t *testing.T) {
	t.Run("should set family_id", func(t *testing.T) {
		opts := FilterOptions{}
		WithFamilyId("PT")(&opts)
		assert.Equal(t, "PT", *opts.familyId)
	})
}

func TestGetFamilyId(t *testing.T) {
	t.Run("should get family_id", func(t *testing.T) {
		familyId := "PT"
		opts := FilterOptions{
			familyId: &familyId,
		}
		assert.Equal(t, "PT", *GetFamilyId(opts))
	})
}

func TestWithParentId(t *testing.T) {
	t.Run("should set parent_id", func(t *testing.T) {
		opts := FilterOptions{}
		WithParentId("PT")(&opts)
		assert.Equal(t, "PT", *opts.parentId)
	})
}

func TestGetParentId(t *testing.T) {
	t.Run("should get parent_id", func(t *testing.T) {
		parentId := "PT"
		opts := FilterOptions{
			parentId: &parentId,
		}
		assert.Equal(t, "PT", *GetParentId(opts))
	})
}
