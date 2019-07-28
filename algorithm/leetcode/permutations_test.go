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
	nums := []int{1, 2, 3, 4}
	res := permute(nums)
	fmt.Println(res)
}

func TestPermute2(t *testing.T) {
	nums := []int{1, 1, 2, 2}
	res := permuteUnique(nums)
	fmt.Println(res)
}
