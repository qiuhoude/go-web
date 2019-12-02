package leetcode

import (
	"sort"
	"testing"
)

//435. 无重叠区间 https://leetcode-cn.com/problems/non-overlapping-intervals/

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	// 求最多有多少个子区域不重叠 ,然后 len(intervals) - 最多个数
	sort.Slice(intervals, func(i, j int) bool { // 结束位置升序
		return intervals[i][1] < intervals[j][1]
	})
	ans := 1
	dp := make([]int, n) // 状态表存储的是 ,到该下标 最大不重复的个数
	dp[0] = 1            // 最少也有一个
	for i := 1; i < n; i++ {
		maxCnt := 0
		for j := i - 1; j >= 0; j-- { // 当前位置 与前的区域进行对比看是否有有重叠
			if !isOverlapping(intervals[j], intervals[i]) {
				maxCnt = dp[j]
				break
			}
		}
		dp[i] = max(maxCnt+1, dp[i-1])
		ans = max(ans, dp[i])
	}
	/*lastEnd := intervals[0][1]
	for i := range intervals {
		if intervals[i][0] < lastEnd { //覆盖了
			continue
		}
		lastEnd = intervals[i][0]
		ans ++
	}*/

	return n - ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 判断两个区间是否有重复
func isOverlapping(rg1, rg2 []int) bool {

	return (rg1[0] < rg2[0] && rg2[0] < rg1[1]) ||
		(rg1[0] < rg2[1] && rg2[1] < rg1[1]) ||
		(rg2[0] < rg1[0] && rg1[0] < rg2[1]) ||
		(rg2[0] < rg1[1] && rg1[1] < rg2[1]) ||
		rg1[0] == rg2[0] && rg1[1] == rg2[1]
}

func TestEraseOverlapIntervals(t *testing.T) {
	arg := [][]int{{1, 2}, {2, 3}}
	ret := eraseOverlapIntervals(arg)
	t.Log(ret)
}
func TestOverlapping(t *testing.T) {

	tests := []struct {
		name string
		arg1 []int
		arg2 []int
		want bool
	}{
		{"[1,2],[1,3]", []int{1, 2}, []int{1, 3}, true},
		{"[1,2],[3,4]", []int{1, 2}, []int{3, 4}, false},
		{"[4,8],[3,4]", []int{4, 8}, []int{3, 4}, false},
		{"[4,8],[3,6]", []int{4, 8}, []int{3, 6}, true},
		{"[1,2],[1,2]", []int{1, 2}, []int{1, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOverlapping(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("isOverlapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
