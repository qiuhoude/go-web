package graph

import (
	"fmt"
	"github.com/qiuhoude/go-web/algorithm/datastructures/queue"
	"github.com/qiuhoude/go-web/algorithm/datastructures/stack"
)

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

type edge struct {
	sid    int // 起点id
	tid    int // 目标id
	weight int // 权重
}

func newEdge(sid, tid, weight int) *edge {
	return &edge{
		sid:    sid,
		tid:    tid,
		weight: weight,
	}
}

// 距离起点的距离
type vertexDis struct {
	id  int
	dis int // 距离起点的距离
}

func (v *vertexDis) String() string {
	return fmt.Sprintf("(%v, %v)", v.id, v.dis)
}

func newVertexDis(id, dis int) *vertexDis {
	return &vertexDis{id: id, dis: dis}
}

// 有向图
type Digraph struct {
	adj   map[int]*vertex
	edges map[int][]*edge // [顶点起点id]
}

func NewGraph() *Digraph {
	return &Digraph{
		adj:   make(map[int]*vertex),
		edges: make(map[int][]*edge),
	}
}

func (g *Digraph) ContainsVertex(id int) bool {
	_, ok := g.adj[id]
	return ok
}

func (g *Digraph) ContainsEdge(sid, tid int) bool {
	if eds, ok := g.edges[sid]; ok {
		for i := range eds {
			if tid == eds[i].tid {
				return true
			}
		}
	}
	return false
}

func (g *Digraph) AddVertex(id int, v interface{}) {
	if !g.ContainsVertex(id) {
		g.adj[id] = newVertex(id, v)
	}
}

// 添加单向边
func (g *Digraph) AddEdge(sid, tid, weight int) {
	if g.ContainsVertex(sid) && g.ContainsVertex(tid) && sid != tid {
		if !g.ContainsEdge(sid, tid) {
			eds := g.edges[sid]
			eds = append(eds, newEdge(sid, tid, weight))
			g.edges[sid] = eds
		}
	}
}

func (g *Digraph) GetVertexData(id int) interface{} {
	if ret, ok := g.adj[id]; ok {
		return ret.v
	}
	return nil
}

// 添加双向边
func (g *Digraph) AddDuplexEdge(sid, tid, weight int) {
	g.AddEdge(sid, tid, weight)
	g.AddEdge(tid, sid, weight)
}

/*
最短路径算法还有 Bellford 算法、Floyd 算法等

Dijkstra 最短路径算法 搜索 时间 O(E*logV) logV是优先队列的是复杂度
*/
func (g *Digraph) DijkstraSearch(sid, tid int) []int {
	if sid == tid { // 同一个点
		return nil
	}
	// 两个点必须存在
	if !g.ContainsVertex(sid) || !g.ContainsVertex(tid) {
		return nil
	}
	vertexDisMap := make(map[int]*vertexDis)
	for id := range g.adj {
		vertexDisMap[id] = newVertexDis(id, 0x7fffffff) //初始值 最大距离
	}

	n := len(g.adj)
	prev := make(map[int]int, n)     // precurosor记录前驱节点,用于反向找path
	visited := make(map[int]bool, n) // 访问标记
	visited[sid] = true
	que := queue.NewPriorityQueue(compareEdge)
	vertexDisMap[sid].dis = 0 // 起点到起点距离是0
	que.Enqueue(vertexDisMap[sid])
	isFound := false

	for !que.IsEmpty() {
		minVertex := que.Dequeue().(*vertexDis)
		if minVertex.id == tid {
			isFound = true
			break
		}
		// 取出一条 minVertex 相连的边
		for _, edge := range g.edges[minVertex.id] {
			// 找到一条到 nextVertex 更短的路径
			nextVertex := vertexDisMap[edge.tid] // minVertex-->nextVertex
			if nextVertex.dis > minVertex.dis+edge.weight {
				nextVertex.dis = minVertex.dis + edge.weight
				prev[nextVertex.id] = minVertex.id
				if !visited[nextVertex.id] {
					visited[nextVertex.id] = true
					que.Enqueue(vertexDisMap[nextVertex.id])
				}
			}
		}
	}
	fmt.Println(vertexDisMap)
	if isFound {
		cur := tid
		st := stack.NewArrayStack()

		for i := 0; i < len(prev); i++ { //反向查找,就是路径
			if _, ok := prev[cur]; !ok {
				break
			}
			st.Push(cur)
			cur = prev[cur]
		}
		st.Push(sid)

		ret := make([]int, st.Len())

		for i := 0; !st.IsEmpty(); i++ {
			ret[i] = st.Pop().(int)
		}
		return ret
	}
	return nil
}

// 边的比较
func compareEdge(v1, v2 interface{}) int {
	ed1 := v1.(*vertexDis)
	ed2 := v2.(*vertexDis)
	if ed1.dis > ed2.dis {
		return 1
	} else if ed1.dis < ed2.dis {
		return -1
	} else {
		return 0
	}
}
