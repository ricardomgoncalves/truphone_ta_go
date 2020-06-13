package queue

import (
	"testing"

	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/stretchr/testify/assert"
)

func TestNewMemberQueue(t *testing.T) {
	t.Run("should not return nil", func(t *testing.T) {
		assert.NotNil(t, NewMemberQueue())
	})
}

func TestMemberQueue_Enqueue(t *testing.T) {
	t.Run("should enqueue new node", func(t *testing.T) {
		queue := NewMemberQueue()
		queue.Enqueue(family.Member{})
		assert.Equal(t, 1, len(queue.items))

		queue.Enqueue(family.Member{})
		assert.Equal(t, 2, len(queue.items))
	})
}

func TestMemberQueue_Dequeue(t *testing.T) {
	t.Run("should dequeue node", func(t *testing.T) {
		queue := NewMemberQueue()
		queue.Enqueue(family.Member{Id: "id"})
		assert.Equal(t, 1, len(queue.items))

		queue.Enqueue(family.Member{Id: "id2"})
		assert.Equal(t, 2, len(queue.items))

		member := queue.Dequeue()
		assert.Equal(t, 1, len(queue.items))
		assert.Equal(t, family.Member{Id: "id"}, member)

		member = queue.Dequeue()
		assert.Equal(t, 0, len(queue.items))
		assert.Equal(t, family.Member{Id: "id2"}, member)
	})
}

func TestMemberQueue_IsEmpty(t *testing.T) {
	t.Run("should check if queue is empty", func(t *testing.T) {
		queue := NewMemberQueue()
		assert.True(t, queue.IsEmpty())

		queue.Enqueue(family.Member{Id: "id"})
		assert.False(t, queue.IsEmpty())
	})
}
