package leetcode

import (
	"github.com/bmizerany/assert"
	"testing"
)

// leetcode 307号问题
// https://leetcode.com/problems/range-sum-query-mutable/description/
// leetcode 303. 区域和检索 - 数组不可变
// https://leetcode-cn.com/problems/range-sum-query-immutable/

type NumArray struct {
	nums []int
	tree []int
}

func Constructor1(nums []int) NumArray {
	n := len(nums)
	ret := NumArray{
		nums: nums,
		tree: make([]int, 4*n),
	}
	if n > 0 {
		ret.buildTree(0, 0, n-1)
	}
	return ret
}

func (this *NumArray) query(i, l, r, ql, qr int) int {
	if r < 0 {
		return -1
	}
	if l == ql && r == qr {
		return this.tree[i]
	}
	// 分3段
	mid := l + (r-l)>>1
	lci := i*2 + 1 // left child index
	rci := i*2 + 2 // right child index
	if qr <= mid { // 区域范围再左边
		return this.query(lci, l, mid, ql, qr)
	} else if ql >= mid+1 { // 区域范围全在右边
		return this.query(rci, mid+1, r, ql, qr)
	} else { // 两边个一半
		sumL := this.query(lci, l, mid, ql, mid)
		sumR := this.query(rci, mid+1, r, mid+1, qr)
		return sumL + sumR
	}

}

func (this *NumArray) Update(i int, val int) {
	n := len(this.nums)
	if i < 0 || i >= n {
		return
	}
	this.nums[i] = val // 修改原数组
	this.set(i, val, 0, 0, len(this.nums)-1)

}

func (this *NumArray) set(ni int, val int, i, li, ri int) {
	if li == ri {
		this.tree[i] = val
		return
	}
	mid := li + (ri-li)>>1
	lci := i*2 + 1 // left child index
	rci := i*2 + 2 // right child index
	if ni <= mid {
		this.set(ni, val, lci, li, mid)
	} else if ni > mid {
		this.set(ni, val, rci, mid+1, ri)
	}
	this.tree[i] = this.tree[lci] + this.tree[rci]
}

func (this *NumArray) buildTree(i, li, ri int) {
	if li == ri {
		this.tree[i] = this.nums[li]
		return
	}
	mid := li + (ri-li)>>1

	lci := i*2 + 1 // left child index
	rci := i*2 + 2 // right child index
	this.buildTree(lci, li, mid)
	this.buildTree(rci, mid+1, ri)
	this.tree[i] = this.tree[lci] + this.tree[rci] // 中间位置的线段树

}

func (this *NumArray) SumRange(i int, j int) int {
	if i > j {
		return -1
	}
	return this.query(0, 0, len(this.nums)-1, i, j)
}

func TestSumRange(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor1(nums)
	assert.Equal(t, obj.SumRange(0, 2), 1)
	assert.Equal(t, obj.SumRange(2, 5), -1)
	assert.Equal(t, obj.SumRange(0, 5), -3)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
