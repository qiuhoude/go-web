package topology

import (
	"testing"
)

func buildGraph() *Graph {
	g := NewGraph()
	g.AddVertex(1, "A")
	g.AddVertex(2, "B")
	g.AddVertex(3, "C")
	g.AddVertex(4, "D")
	//g.AddVertex(5, "E")
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(4, 2)
	//g.AddEdge(3, 1)
	return g
}

func Test_topoSortKahn(t *testing.T) {
	g := buildGraph()
	topSort := topoSortKahn(g)
	t.Log(topSort)
}

func Test_topoSortDFS(t *testing.T) {
	g := buildGraph()
	topoSortDFS(g)
}
