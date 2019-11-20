package graph

import (
	"fmt"
	"testing"
)

func buildDiGraph() *Digraph {
	g := NewGraph()
	g.AddVertex(1, "莫斯科")
	g.AddVertex(2, "华盛顿")
	g.AddVertex(3, "伦敦")
	g.AddVertex(4, "圣彼得堡")
	g.AddVertex(5, "哥本哈根")
	g.AddVertex(6, "里斯本")
	g.AddVertex(7, "新奥尔良")
	g.AddVertex(8, "阿姆斯特丹")
	g.AddVertex(9, "布鲁塞尔")
	g.AddVertex(10, "巴黎")
	// 添加边
	g.AddDuplexEdge(1, 4, 1)
	g.AddDuplexEdge(1, 5, 1)
	g.AddDuplexEdge(2, 7, 1)
	g.AddDuplexEdge(3, 8, 1)
	g.AddDuplexEdge(2, 6, 1)
	g.AddDuplexEdge(3, 9, 1)
	g.AddDuplexEdge(4, 10, 1)
	g.AddDuplexEdge(5, 10, 1)
	g.AddDuplexEdge(6, 10, 1)
	g.AddDuplexEdge(7, 10, 1)
	g.AddDuplexEdge(8, 10, 1)
	g.AddDuplexEdge(9, 10, 1)
	return g
}

func TestGraphAdjTab_DijkstraSearch(t *testing.T) {
	g := buildDiGraph()
	path := g.DijkstraSearch(2, 3)
	for _, v := range path {
		fmt.Printf("(%v)%v->", v, g.GetVertexData(v))
	}
	fmt.Println()
}
