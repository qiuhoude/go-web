package leetcode

import (
	"sort"
)

// 15. 三数之和 https://leetcode-cn.com/problems/3sum/

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
func threeSum(nums []int) [][]int {
	// 思路: 1.排序,
	// 2. 使用三指针 i l r

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j] // 正序
	})

	var res [][]int
	length := len(nums)
	if length < 3 || nums[0] > 0 {
		// 无解情况
		return nil
	}

	for i := 0; i < length && nums[i] <= 0; i++ {
		l := i + 1
		r := length - 1
		if i > 0 && nums[i] == nums[i-1] { // 重复的去掉
			continue
		}
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 { // 找到目的
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] { // 去重
					l++
				}
				for l < r && nums[r] == nums[r-1] { // 去重
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			}
		}
	}

	return res
}
