package leetcode

// leetcode 307号问题
// https://leetcode.com/problems/range-sum-query-mutable/description/
type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	n := NumArray{}
	n.sum = make([]int, len(nums)+1)
	n.sum[0] = 0
	for i := 1; i < len(nums)+1; i++ {
		n.sum[i] = nums[i-1] + n.sum[i-1]
	}
	return n
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sum[j+1] - this.sum[i]
}
