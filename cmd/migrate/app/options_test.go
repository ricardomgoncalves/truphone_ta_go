package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOptions(t *testing.T) {
	t.Run("should return empty options", func(t *testing.T) {
		opts := NewOptions()
		assert.Equal(t, &Options{}, opts)
	})
}