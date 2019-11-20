package graph

import (
	"fmt"
	"github.com/qiuhoude/go-web/algorithm/datastructures/queue"
	"github.com/qiuhoude/go-web/algorithm/datastructures/stack"
)

/*
基本概念
方向维度: 无向图 有向图
权重维度: 带权图 非带权图
稀疏维度: 稠密图 稀疏图

图的存储方式
1. 邻接矩阵
2. 邻接表(adjacency table) , 优化后 可以用跳表 或 树结构代替普通链表加快查询

搜索算法
1. BST(Breadth-First-Search) 广度优先 借用queue
时间 O(V+E):V 表示顶点的个数，E 表示边的个数
空间 O(V) visited、prev


2. DST(Depth-First-Search) 深度优先 借用stack
时间 O(E) 每条边最多会被访问两次，一次是遍历，一次是回退
空间 O(V) 主要消耗在 visited、prev 数组和递归调用栈,但最大深度不会超过V个节点
*/

// 图的节点
type node struct {
	id     int           // 节点id
	degree map[int]*node // 连接的节点
	v      interface{}   // 节点存储数据
}

func newGraphNode(id int, v interface{}) *node {
	return &node{
		id:     id,
		degree: make(map[int]*node),
		v:      v,
	}
}

// 邻接表的无向图
type GraphAdjTab struct {
	adj map[int]*node
	// 此处可以有个逆邻接表
}

func NewGraphAdjTab() *GraphAdjTab {
	return &GraphAdjTab{adj: make(map[int]*node)}
}

// 向无向图中添加顶点
func (g *GraphAdjTab) AddVertex(id int, v interface{}) {
	if c, ok := g.adj[id]; ok {
		c.v = v
	} else {
		g.adj[id] = newGraphNode(id, v)
	}
}

// 向无向图中添加边, 点必须存在
func (g *GraphAdjTab) AddEdge(sid, tid int) bool {
	if !g.containVertex(sid) || !g.containVertex(tid) {
		return false
	}

	sNode := g.adj[sid]
	tNode := g.adj[tid]
	sNode.degree[tid] = tNode
	tNode.degree[sid] = sNode
	return true
}

// 是否包含顶点
func (g *GraphAdjTab) containVertex(id int) bool {
	_, ok := g.adj[id]
	return ok
}

// 根据id获取节点数据
func (g *GraphAdjTab) GetVertexData(id int) interface{} {
	if ret, ok := g.adj[id]; ok {
		return ret.v
	}
	return nil
}

// 广度优先路径搜索, sid 到 tid的路径 返回路线id
func (g *GraphAdjTab) BSTSearch(sid, tid int) []int {
	if sid == tid { // 同一个点
		return nil
	}
	// 两个点必须存在
	if !g.containVertex(sid) || !g.containVertex(tid) {
		return nil
	}

	n := len(g.adj)
	prev := make(map[int]int, n)
	visited := make(map[int]bool, n) // 访问标记
	visited[sid] = true              // 起点已经访问
	que := queue.NewArrayQueue(n)
	que.Enqueue(sid)
	isFound := false
out:
	for que.Len() > 0 && !isFound {
		e := que.Dequeue()
		id := e.(int)
		node := g.adj[id]
		for idk := range node.degree {
			if !visited[idk] { // 没有被访问
				prev[idk] = id
				//fmt.Printf("%v->%v ;", id, idk)
				if idk == tid { // 已经找到
					isFound = true
					break out
				}
				visited[idk] = true // 标记已访问
				que.Enqueue(idk)
			}
		}
		//fmt.Println()
	}
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

// 深度优先, sid 到 tid的路径 返回路线id
func (g *GraphAdjTab) DSTSearch(sid, tid int) []int {
	if sid == tid { // 同一个点
		return nil
	}
	// 两个点必须存在
	if !g.containVertex(sid) || !g.containVertex(tid) {
		return nil
	}
	n := len(g.adj)
	prev := make(map[int]int, n)
	visited := make(map[int]bool, n) // 访问标记
	visited[sid] = true              // 起点已经访问
	st := stack.NewArrayStack()
	isFound := false
	cur := g.adj[sid]
	st.Push(sid)
out:
	for !isFound && cur != nil {
		flag := false // 是否还能继续往下找
		for idk := range cur.degree {
			if !visited[idk] {
				st.Push(idk)
				visited[idk] = true //标记以访问
				prev[idk] = cur.id
				fmt.Printf("forward %v->%v \n", cur.id, idk)
				if idk == tid { // 已经找到
					isFound = true
					break out
				}
				cur = g.adj[idk] // 进行下一层
				flag = true
				break
			}
		}
		if !flag && !st.IsEmpty() {
			// 退栈
			id := st.Pop().(int)
			fmt.Printf("back %v->%v \n", cur.id, id)
			cur = g.adj[id]
		}
	}
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
