// 并查集
package unionfind

import "errors"

type Uf interface {
	GetSize() int
	IsConnected(p, q int) bool    // p,q两个下标是是否连接
	UnionElements(p, q int) error // p,q下标进行合并
}

type UnionFind struct {
	parent []int
}

var ArgumentErr = errors.New("index is out of bound")

func NewUnionFind(size int) *UnionFind {
	p := make([]int, size)
	for i := 0; i < size; i++ {
		p[i] = i
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
		p = u.parent[p]
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
	// 将p的根节点指向q的根节点
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
