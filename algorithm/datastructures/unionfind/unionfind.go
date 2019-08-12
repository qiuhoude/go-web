// 并查集
package unionfind

import "errors"

type Uf interface {
	GetSize() int
	IsConnected(p, q int) bool    // p,q两个下标是是否连接
	UnionElements(p, q int) error // p,q下标进行合并
}

type UnionFind struct {
	parent []int // parent[i]表示第i个元素所指向的父节点
	rank   []int //rank[i]表示以i为根的集合所表示的树的层数,在路径压缩之后不表示层数而是一个分数
}

var ArgumentErr = errors.New("index is out of bound")

func NewUnionFind(size int) *UnionFind {
	p := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		p[i] = i
		rank[i] = 1
	}
	return &UnionFind{parent: p}
}

func (u *UnionFind) GetSize() int {
	return len(u.parent)
}

// 查找过程, 查找元素p所对应的集合编号
// O(h)复杂度, h为树的高度
func (u *UnionFind) find(p int) (int, error) {
	if p < 0 || p > u.GetSize() {
		return 0, ArgumentErr
	}
	// 不断去寻找自己的父节点, 直到到达根节点
	// 根节点的特点: parent[p] == p
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]] // 将自己的父节点 指向爷节点
		p = u.parent[p]                     // 直接从爷节点开始往上
	}

	return p, nil
}

func (u *UnionFind) UnionElements(p, q int) error {
	pRoot, err := u.find(p)
	if err != nil {
		return err
	}
	qRoot, err := u.find(p)
	if err != nil {
		return err
	}
	if qRoot == pRoot {
		return nil
	}

	// 根据两个元素所在树的rank不同判断合并方向
	// 将rank低的集合合并到rank高的集合上
	if u.rank[pRoot] < u.rank[qRoot] {
		u.parent[pRoot] = qRoot
	} else if u.rank[pRoot] > u.rank[qRoot] {
		u.parent[qRoot] = pRoot
	} else { // rank 相等
		u.parent[pRoot] = qRoot
		u.rank[qRoot] += 1
	}
	u.parent[pRoot] = qRoot
	return nil
}

func (u *UnionFind) IsConnected(p, q int) bool {
	pRoot, err := u.find(p)
	if err != nil {
		return false
	}
	qRoot, err := u.find(p)
	if err != nil {
		return false
	}
	return pRoot == qRoot
}
