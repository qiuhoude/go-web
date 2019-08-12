package unionfind

// 基于size的优化,将集
type UnionFindOpSize struct {
	parent []int
	sz     []int
}

func NewUnionFindOpSize(size int) *UnionFindOpSize {
	p := make([]int, size)
	sz := make([]int, size)
	for i := 0; i < size; i++ {
		p[i] = i
		sz[i] = 1
	}
	ret := &UnionFindOpSize{sz: sz}
	ret.parent = p
	return ret
}

func (u *UnionFindOpSize) GetSize() int {
	return len(u.parent)
}

// 查找过程, 查找元素p所对应的集合编号
// O(h)复杂度, h为树的高度
func (u *UnionFindOpSize) find(p int) (int, error) {
	if p < 0 || p > u.GetSize() {
		return 0, ArgumentErr
	}
	// 不断去寻找自己的父节点, 直到到达根节点
	// 根节点的特点: parent[p] == p
	for p != u.parent[p] {
		p = u.parent[p]
	}

	return p, nil

	// 递归算法
	//if p != u.parent[p] {
	//	u.parent[p], _ = u.find(u.parent[p]) // 将找到的父节点的父节点复制给父节点
	//}
	//return u.parent[p], nil
}

func (u *UnionFindOpSize) IsConnected(p, q int) bool {
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

func (u *UnionFindOpSize) UnionElements(p, q int) error {
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
	// 根据两个元素所在树的元素个数不同判断合并方向
	// 将元素个数少的集合合并到元素个数多的集合上
	if u.sz[pRoot] < u.sz[qRoot] {
		u.parent[pRoot] = qRoot
		u.sz[qRoot] += u.sz[pRoot]
	} else {
		u.parent[qRoot] = pRoot
		u.sz[pRoot] += u.sz[qRoot]
	}
	return nil
}
