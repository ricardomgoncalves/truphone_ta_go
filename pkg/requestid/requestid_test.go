package requestid

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRequestId(t *testing.T) {
	t.Run("should return not ok on empty context", func(t *testing.T) {
		id, ok := GetRequestId(context.Background())
		assert.Equal(t, "", id)
		assert.False(t, ok)
	})
	t.Run("should return ok with id", func(t *testing.T) {
		id, ok := GetRequestId(context.WithValue(context.Background(), requestIdKey, "id"))
		assert.Equal(t, "id", id)
		assert.True(t, ok)
	})
}

func TestWithRequestId(t *testing.T) {
	t.Run("should set id on context", func(t *testing.T) {
		ctx := WithRequestId(context.Background(), "id")
		id, ok := ctx.Value(requestIdKey).(string)
		assert.Equal(t, "id", id)
		assert.True(t, ok)
	})
}
