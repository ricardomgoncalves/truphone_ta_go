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