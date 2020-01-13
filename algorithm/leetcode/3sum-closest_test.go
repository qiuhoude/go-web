package leetcode

import "sort"

// 16. 最接近的三数之和 https://leetcode-cn.com/problems/3sum-closest/

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。
返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

*/

func threeSumClosest(nums []int, target int) int {
	// 思路: 1. 排序
	// 2. 也是使用三指针
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j] // 正序
	})

	length := len(nums)
	if length < 3 {
		return 0
	}
	ret := nums[0] + nums[1] + nums[2]

	for i := 0; i < length; i++ {
		l := i + 1
		r := length - 1
		if i > 0 && nums[i] == nums[i-1] { // 重复的去掉
			continue
		}
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(target-ret) == 0 { // 找到目的
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
	return ret
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
