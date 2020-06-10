package family

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewFamilyWithId(t *testing.T) {
	t.Run("should return family with id", func(t *testing.T) {
		id, err := uuid.NewUUID()
		require.Nil(t, err)

		family := NewFamilyWithId(id)
		assert.Equal(t, id, family.Id)
	})
}