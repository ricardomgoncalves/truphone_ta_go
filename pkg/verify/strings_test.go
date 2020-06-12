package verify

import (
	"testing"

	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestStringLength(t *testing.T) {
	t.Run("should return error on only spaces string", func(t *testing.T) {
		assert.Equal(t, errors.New("name should have a character different of a space"), StringLength("     ", 3, 30))
	})
	t.Run("should return error length less than 3", func(t *testing.T) {
		assert.Equal(t, errors.New("name should be at least 3 characters"), StringLength("aa", 3, 30))
		assert.Equal(t, errors.New("name should be at least 5 characters"), StringLength("aa", 5, 30))
	})
	t.Run("should return error length greater than 30", func(t *testing.T) {
		assert.Equal(t, errors.New("name should be at maximum 30 characters"), StringLength("this name is too long be a part of this", 3, 30))
		assert.Equal(t, errors.New("name should be at maximum 2 characters"), StringLength("this name is too long be a part of this", 1, 2))
	})
	t.Run("should return no error", func(t *testing.T) {
		assert.Equal(t, nil, StringLength("this name", 3, 30))
	})
}
