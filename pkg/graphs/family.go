package graphs

import (
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/queue"
)

type FamilyGraph struct {
	nodes []family.Member
	edges map[string][]family.Member
}

func NewFamilyGraph() *FamilyGraph {
	return &FamilyGraph{
		nodes: make([]family.Member, 0),
	}
}

func (g *FamilyGraph) AddNode(n family.Member) {
	g.nodes = append(g.nodes, n)
}

func (g *FamilyGraph) AddEdge(from string, to family.Member) {
	if g.edges == nil {
		g.edges = make(map[string][]family.Member)
	}
	g.edges[from] = append(g.edges[from], to)
}

func (g *FamilyGraph) Traverse(f func(children []family.Member)) {
	q := queue.NewMemberQueue()
	visited := make(map[string]struct{})

	n := g.nodes[0]
	q.Enqueue(n)

	for {
		if q.IsEmpty() {
			break
		}

		node := q.Dequeue()
		visited[node.Id] = struct{}{}
		children := g.edges[node.Id]

		for i := 0; i < len(children); i++ {
			child := children[i]
			if _, ok := visited[child.Id]; !ok {
				q.Enqueue(child)
				visited[child.Id] = struct{}{}
			}
		}

		if f != nil {
			if len(children) != 0 {
				f(children)
			}
		}
	}
}
