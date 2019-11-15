package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

// 46. 全排列
//https://leetcode-cn.com/problems/permutations/

// 思路: 使用递归回溯法
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	if len(nums) == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}

	var result [][]int
	for index, value := range nums {
		var numsCopy = make([]int, len(nums))
		copy(numsCopy, nums)
		// 将numsCopy中index这个元素给剔除掉赋值给numsSubOne
		numsSubOne := append(numsCopy[:index], numsCopy[index+1:]...)
		valueSlice := []int{value}
		newSubSlice := permute(numsSubOne)
		for _, newValue := range newSubSlice {
			result = append(result, append(valueSlice, newValue...))
		}
	}
	return result
}

// 47  全排列2,nums中有重复的元素
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	if len(nums) == 2 {
		if nums[0] == nums[1] { // 与前一个数相等 只返回一种可能性
			return [][]int{{nums[0], nums[1]}}
		} else {
			return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
		}
	}
	// 先排序
	sort.Ints(nums)
	var result [][]int
	for index, value := range nums {
		var numsCopy = make([]int, len(nums))
		copy(numsCopy, nums)
		if index > 0 && numsCopy[index] == numsCopy[index-1] {
			continue
		}
		// 将numsCopy中index这个元素给剔除掉赋值给numsSubOne
		numsSubOne := append(numsCopy[:index], numsCopy[index+1:]...)
		valueSlice := []int{value}
		newSubSlice := permuteUnique(numsSubOne)
		for _, newValue := range newSubSlice {
			result = append(result, append(valueSlice, newValue...))
		}
	}
	return result
}

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}

func TestPermute2(t *testing.T) {
	nums := []int{1, 1, 1, 2}
	//res := permuteUnique(nums)
	res := permuteUnique2(nums)
	fmt.Println(res)
}

func permuteUnique2(nums []int) [][]int {
	var res [][]int
	// 先排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	used := make([]bool, len(nums))
	var stack []int
	permuteHelper(0, nums, used, stack, &res)
	return res
}

func permuteHelper(index int, nums []int, used []bool, stack []int, res *[][]int) {
	n := len(nums)
	if index == n {
		// 出口已经找到
		arr := make([]int, n)
		copy(arr, stack)
		*res = append(*res, arr)
		return
	}
	for i := 0; i < n; i++ {
		if !used[i] {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] {
				continue
			}
			used[i] = true
			permuteHelper(index+1, nums, used, append(stack, nums[i]), res)
			used[i] = false
		}
	}

}
