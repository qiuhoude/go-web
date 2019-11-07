package graph

import (
	"fmt"
	"testing"
)

func buildGraph() *GraphAdjTab {
	g := NewGraphAdjTab()
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
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 6)
	g.AddEdge(2, 7)
	g.AddEdge(3, 8)
	g.AddEdge(3, 9)
	g.AddEdge(4, 10)
	g.AddEdge(5, 10)
	g.AddEdge(6, 10)
	g.AddEdge(7, 10)
	g.AddEdge(8, 10)
	g.AddEdge(9, 10)
	return g
}
func TestGraphAdjTab_BSTSearch(t *testing.T) {
	g := buildGraph()
	path := g.BSTSearch(2, 3)
	fmt.Printf("(%d)%v", 2, g.GetVertexData(2))
	for _, v := range path {
		fmt.Printf("->(%d)%v", v, g.GetVertexData(v))
	}
	fmt.Println()
}

func TestGraphAdjTab_DSTSearch(t *testing.T) {
	g := buildGraph()
	path := g.DSTSearch(7, 9)
	for _, v := range path {
		fmt.Printf("->%v", g.GetVertexData(v))
	}
	fmt.Println()
}
