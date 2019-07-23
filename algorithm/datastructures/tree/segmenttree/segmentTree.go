package segmenttree

import (
	"fmt"
	"strings"
)

type MergeFunc func(l, r interface{}) interface{}

type SegmentTree struct {
	data  []interface{} // 线段数存储的数据
	tree  []interface{}
	merge MergeFunc
}

func NewSegmentTree(data []interface{}, mf MergeFunc) *SegmentTree {
	if data == nil || len(data) == 0 {
		return nil
	}
	ret := &SegmentTree{
		data:  data,
		merge: mf,
	}
	// 线段数是非满二叉树, 通过等比数列求和公式可得 2^n - 1 ,4n就可以容纳所有线段树
	size := 4 * len(data)
	ret.tree = make([]interface{}, size, size) //直接申请这么多的空间
	ret.buildSegmentTree(0, 0, len(data)-1)
	return ret
}

// 构建线段树
func (st *SegmentTree) buildSegmentTree(treeIndex, l, r int) {
	if l == r {
		st.tree[treeIndex] = st.data[l]
		return
	}
	lTreeIndex := leftChild(treeIndex)
	rTreeIndex := rightChild(treeIndex)

	mid := l + (r-l)/2
	st.buildSegmentTree(lTreeIndex, l, mid) // 构建左边的
	st.buildSegmentTree(rTreeIndex, mid+1, r)

	st.tree[treeIndex] = st.merge(st.tree[lTreeIndex], st.tree[rTreeIndex])
}

func (st *SegmentTree) Size() int {
	if st == nil {
		return 0
	}
	return len(st.data)
}

// 进行范围查询
func (st *SegmentTree) Query(queryL, queryR int) interface{} {
	if queryL < 0 || queryL >= len(st.data) ||
		queryR < 0 || queryR >= len(st.data) || queryL > queryR {
		return nil
	}
	return st.query(0, 0, st.Size()-1, queryL, queryR)
}

// 在以treeIndex为根的线段树中[l...r]的范围里，搜索区间[queryL...queryR]的值
func (st *SegmentTree) query(treeIndex, l, r, queryL, queryR int) interface{} {
	if l == queryL && r == queryR { //直接找到到了
		return st.tree[treeIndex]
	}
	lTreeIndex := leftChild(treeIndex)
	rTreeIndex := rightChild(treeIndex)
	mid := l + (r-l)/2
	if queryL >= mid+1 { // 区间全部落在右边
		return st.query(rTreeIndex, mid+1, r, queryL, queryR)
	} else if queryR <= mid { //区间全部落在左边
		return st.query(lTreeIndex, l, mid, queryL, queryR)
	} else { // 左右各占一半
		leftRes := st.query(lTreeIndex, l, mid, queryL, mid)
		rightRes := st.query(rTreeIndex, mid+1, r, mid+1, queryR)
		return st.merge(rightRes, leftRes)
	}
}

// 左孩子下标(和二叉堆一样)
func leftChild(index int) int {
	return index*2 + 1
}

// 右孩子下标
func rightChild(index int) int {
	return index*2 + 2
}

// 获取当前层级

func (st *SegmentTree) String() string {
	sb := strings.Builder{}
	sb.WriteString("[")
	for i := 0; i < len(st.tree); i++ {
		if st.tree[i] == nil {
			sb.WriteString("nil")
		} else {
			sb.WriteString(fmt.Sprintf("%v", st.tree[i]))
		}
		if i != len(st.tree)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func depth(index int) int {
	i := 1
	cnt := 0
	for index >= i {
		i = i << 1
		cnt++
	}
	return cnt
}
