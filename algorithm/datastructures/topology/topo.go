package topology

import (
	"fmt"
	"github.com/qiuhoude/go-web/algorithm/datastructures/queue"
	"github.com/qiuhoude/go-web/algorithm/datastructures/stack"
)

/*

拓扑排序中
1. kahn算法
2. DFS算法

比如 依赖关系 (箭头表示 xx 依赖 xxx ) A->B , B->C ,C->D
求出依赖关系
使用 有向图 的数据结构进行处理

*/

// 顶点
type vertex struct {
	id int
	v  interface{}
}

func (v *vertex) String() string {
	return fmt.Sprintf("(%v, %v)", v.id, v.v)
}

func newVertex(id int, v interface{}) *vertex {
	return &vertex{
		id: id,
		v:  v,
	}
}

// 边
type edge struct {
	sid int // 起点id
	tid int // 目标id
	//weight int // 权重
}

func newEdge(sid, tid int) *edge {
	return &edge{
		sid: sid,
		tid: tid,
	}
}

// 有向图
type Graph struct {
	adj   map[int]*vertex
	edges map[int][]*edge // [顶点起点id]
}

func NewGraph() *Graph {
	return &Graph{
		adj:   make(map[int]*vertex),
		edges: make(map[int][]*edge),
	}
}

// 添加顶点
func (g *Graph) AddVertex(id int, v interface{}) {
	if !g.ContainsVertex(id) {
		g.adj[id] = newVertex(id, v)
	}
}
func (g *Graph) ContainsVertex(id int) bool {
	_, ok := g.adj[id]
	return ok
}

func (g *Graph) ContainsEdge(sid, tid int) bool {
	if eds, ok := g.edges[sid]; ok {
		for i := range eds {
			if tid == eds[i].tid {
				return true
			}
		}
	}
	return false
}

// 添加边
func (g *Graph) AddEdge(sid, tid int) {
	if g.ContainsVertex(sid) && g.ContainsVertex(tid) && sid != tid {
		if !g.ContainsEdge(sid, tid) {
			eds := g.edges[sid]
			eds = append(eds, newEdge(sid, tid))
			g.edges[sid] = eds
		}
	}
}

//拓扑排序中的 Kahn 算法  O(V+E)（V 表示顶点个数，E 表示边的个数）
func topoSortKahn(g *Graph) []interface{} {
	// 统计顶点入度
	inDegree := make(map[int]int)
	for id := range g.adj {
		inDegree[id] = 0
	}
	for id := range g.adj {
		if eds, ok := g.edges[id]; ok {
			for i := range eds {
				inDegree[eds[i].tid]++
			}
		}
	}
	que := queue.NewLinkedQueue()
	for i := range inDegree {
		if inDegree[i] == 0 { // 入度是0的添加进队列
			que.Enqueue(i)
		}
	}
	st := stack.NewArrayStack()
	for !que.IsEmpty() {
		id := que.Dequeue().(int)
		st.Push(g.adj[id].v)
		//fmt.Print("->", g.adj[id])
		for _, v := range g.edges[id] {
			inDegree[v.tid]-- //
			if inDegree[v.tid] == 0 {
				que.Enqueue(v.tid)
			}
		}
	}
	if st.Len() < len(g.adj) {
		fmt.Println("有环图")
	}
	return st.ToSlice()
}

// DFS方式排序 时间复杂度 O(V+E)
func topoSortDFS(g *Graph) {
	// 创建逆邻接表
	inverseAdj := make(map[int][]int, len(g.adj))
	// 填充逆邻接表
	for _, eds := range g.edges {
		for i := range eds {
			e := eds[i]
			inverseAdj[e.tid] = append(inverseAdj[e.tid], e.sid)
		}
	}
	visited := make(map[int]bool) // 访问记录
	for v := range g.adj {
		if !visited[v] {
			visited[v] = true
			dfs(v, inverseAdj, visited, g)
		}
	}
	fmt.Println()
}

func dfs(vertex int, inverseAdj map[int][]int, visited map[int]bool, g *Graph) {
	tids := inverseAdj[vertex]
	for i := 0; i < len(tids); i++ {
		tid := tids[i]
		if visited[tid] {
			continue
		}
		visited[tid] = true
		dfs(tid, inverseAdj, visited, g)
	}
	fmt.Print("->", g.adj[vertex].v)
}
