package graphs

import (
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFamilyGraph(t *testing.T) {
	t.Run("should not get nil", func(t *testing.T) {
		assert.NotNil(t, NewFamilyGraph())
	})
}

func TestFamilyGraph_AddNode(t *testing.T) {
	t.Run("should add node to node slice", func(t *testing.T) {
		graph := NewFamilyGraph()
		graph.AddNode(family.Member{})
		assert.Equal(t, 1, len(graph.nodes))

		graph.AddNode(family.Member{})
		assert.Equal(t, 2, len(graph.nodes))
	})
}

func TestFamilyGraph_AddEdge(t *testing.T) {
	t.Run("should add edge to edge map", func(t *testing.T) {
		graph := NewFamilyGraph()
		graph.AddEdge("1", family.Member{})
		assert.Equal(t, 1, len(graph.edges))

		graph.AddEdge("2", family.Member{})
		assert.Equal(t, 2, len(graph.edges))
	})
}
