package leetcode

import "testing"

//leetcode 300. 最长上升子序列
// https://leetcode-cn.com/problems/longest-increasing-subsequence

/*
给定一个无序的整数数组，找到其中最长上升子序列的长度。
输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
*/

// 思路: 求出 f(i) = j i表示数值的下表,j表示长度
// 当下表为i 时 前面序列的长度
// f(0) = 1 10前面没有比他小的
// f(1) = 1 9 前面没有比他小的
// f(2) = f(1)+1 = 2
// f(3) = max(f(0),f(1))+1
// db[] 保存的就是 下表为i的时最长子序列长度

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	max := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		//比较前面i项最大子序列值
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func lengthOfLIS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// tail 数组的定义：长度为 i + 1 的上升子序列的末尾最小是几
	tail := make([]int, len(nums))
	// 遍历第 1 个数，直接放在有序数组 tail 的开头
	tail[0] = nums[0]
	// end 表示有序数组 tail 的最后一个已经赋值元素的索引
	end := 0
	for i := 1; i < len(nums); i++ {
		// 【逻辑 1】比 tail 数组实际有效的末尾的那个元素还大
		if nums[i] > tail[end] {
			end++
			tail[end] = nums[i]
		} else {
			// 使用二分查找法，在有序数组 tail 中
			// 找到第 1 个大于等于 nums[i] 的元素，尝试让那个元素更小
			left := 0
			right := left
			for left < right {
				// 选左中位数不是偶然，而是有原因的，原因请见 LeetCode 第 35 题题解
				mid := left + (right-left)/2
				if tail[mid] < nums[i] {
					// 中位数肯定不是要找的数，把它写在分支的前面
					left = mid + 1
				} else {
					right = mid
				}
			}
			//走到这里是因为 【逻辑 1】 的反面，因此一定能找到第 1 个大于等于 nums[i] 的元素
			tail[left] = nums[i]
		}
	}
	end++
	return end
}

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"[10,9,2,5,3,7,101,18] => 4", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"[10,9,2,5,3,4] => 3", []int{10, 9, 2, 5, 3, 4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)

			}
		})
	}
}
