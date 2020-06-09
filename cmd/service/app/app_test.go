package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServiceApp(t *testing.T) {
	t.Run("should return empty service app", func(t *testing.T) {
		app := NewServiceApp()
		assert.Equal(t, ServiceApp{}, app)
	})
}
